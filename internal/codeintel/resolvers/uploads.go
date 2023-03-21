package resolvers

import (
	"context"

	"github.com/graph-gophers/graphql-go"

	"github.com/sourcegraph/sourcegraph/internal/gqlutil"
)

type UploadsServiceResolver interface {
	CommitGraph(ctx context.Context, id graphql.ID) (CodeIntelligenceCommitGraphResolver, error)
}

type CodeIntelligenceCommitGraphResolver interface {
	Stale(ctx context.Context) (bool, error)
	UpdatedAt(ctx context.Context) (*gqlutil.DateTime, error)
}