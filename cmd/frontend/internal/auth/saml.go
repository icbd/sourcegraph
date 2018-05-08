package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/internal/db"
	"github.com/sourcegraph/sourcegraph/pkg/actor"
	"github.com/sourcegraph/sourcegraph/pkg/conf"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

// SAML App creation vars
var samlProvider = conf.AuthSAML()

// newSAMLAuthMiddleware returns middlewares for SAML authentication, adding endpoints under the auth
// path prefix to enable the login flow an requiring login for all other endpoints.
//
// 🚨 SECURITY
func newSAMLAuthMiddleware(createCtx context.Context, appURL string) (*Middleware, error) {
	if samlProvider == nil {
		return nil, errors.New("No SAML ID Provider specified")
	}
	if samlProvider.ServiceProviderCertificate == "" {
		return nil, errors.New("No SAML Service Provider certificate")
	}
	if samlProvider.ServiceProviderPrivateKey == "" {
		return nil, errors.New("No SAML Service Provider private key")
	}

	samlSP, err := getSAMLServiceProvider(samlProvider)
	if err != nil {
		return nil, err
	}

	idpID := samlSP.ServiceProvider.IDPMetadata.EntityID
	samlAuthIfNeededMiddleware := func(next http.Handler) http.Handler {
		authedHandler := samlSP.RequireAccount(samlToActorMiddleware(next, idpID))
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Respect already authenticated actor (e.g., via access token).
			if actor.FromContext(r.Context()).IsAuthenticated() {
				next.ServeHTTP(w, r)
				return
			}

			// Otherwise require SAML authentication.
			authedHandler.ServeHTTP(w, r)
		})
	}

	return &Middleware{
		API: func(next http.Handler) http.Handler {
			nextWithSAMLAuth := samlAuthIfNeededMiddleware(next)
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if actor.FromContext(r.Context()).IsAuthenticated() {
					// Request is already authenticated (e.g., by access token).
					next.ServeHTTP(w, r)
					return
				}
				if c, _ := r.Cookie(samlSP.CookieName); c != nil {
					// Try to use cookie to authenticate via SAML.
					nextWithSAMLAuth.ServeHTTP(w, r)
					return
				}
				http.Error(w, "requires authentication", http.StatusUnauthorized)
			})
		},
		App: func(next http.Handler) http.Handler {
			next = samlAuthIfNeededMiddleware(next)
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Handle SAML ACS and metadata endpoints.
				if strings.HasPrefix(r.URL.Path, authURLPrefix+"/saml/") {
					samlSP.ServeHTTP(w, r)
					return
				}
				// Handle all other endpoints
				next.ServeHTTP(w, r)
			})
		},
	}, nil
}

// samlToActorMiddleware translates the SAML session into an Actor and sets it in the request context
// before delegating to its child handler.
func samlToActorMiddleware(h http.Handler, idpID string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actr, err := getActorFromSAML(r, idpID)
		if err != nil {
			log15.Error("Error looking up SAML-authenticated user.", "error", err)
			http.Error(w, "Error looking up SAML-authenticated user. "+couldNotGetUserDescription, http.StatusInternalServerError)
			return
		}
		h.ServeHTTP(w, r.WithContext(actor.WithActor(r.Context(), actr)))
	})
}

// getActorFromSAML translates the SAML session into an Actor.
func getActorFromSAML(r *http.Request, idpID string) (*actor.Actor, error) {
	ctx := r.Context()
	subject := r.Header.Get("X-Saml-Subject") // this header is set by the SAML library after extracting the value from the JWT cookie
	externalID := samlToExternalID(idpID, subject)

	email := r.Header.Get("X-Saml-Email")
	if email == "" && mightBeEmail(subject) {
		email = subject
	}
	login := r.Header.Get("X-Saml-Login")
	if login == "" {
		login = r.Header.Get("X-Saml-Uid")
	}
	displayName := r.Header.Get("X-Saml-DisplayName")
	if displayName == "" {
		displayName = login
	}
	if displayName == "" {
		displayName = email
	}
	if displayName == "" {
		displayName = subject
	}
	if login == "" {
		login = email
	}
	if login == "" {
		return nil, fmt.Errorf("could not create user, because SAML assertion did not contain email attribute statement")
	}
	login, err := NormalizeUsername(login)
	if err != nil {
		return nil, err
	}

	userID, err := createOrUpdateUser(ctx, db.NewUser{
		ExternalProvider: idpID,
		ExternalID:       externalID,
		Username:         login,
		Email:            email,
		DisplayName:      displayName,
		// SAML has no standard way of providing an avatar URL.
	})
	if err != nil {
		return nil, err
	}
	return actor.FromUser(userID), nil
}

func samlToExternalID(idpID, subject string) string {
	return fmt.Sprintf("%s:%s", idpID, subject)
}

func mightBeEmail(s string) bool {
	return strings.Count(s, "@") == 1
}
