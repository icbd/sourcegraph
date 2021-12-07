// Code generated by go-mockgen 1.1.2; DO NOT EDIT.

package runner

import (
	"context"
	"sync"

	definition "github.com/sourcegraph/sourcegraph/internal/database/migration/definition"
)

// MockStore is a mock implementation of the Store interface (from the
// package
// github.com/sourcegraph/sourcegraph/internal/database/migration/runner)
// used for unit testing.
type MockStore struct {
	// DownFunc is an instance of a mock function object controlling the
	// behavior of the method Down.
	DownFunc *StoreDownFunc
	// LockFunc is an instance of a mock function object controlling the
	// behavior of the method Lock.
	LockFunc *StoreLockFunc
	// UpFunc is an instance of a mock function object controlling the
	// behavior of the method Up.
	UpFunc *StoreUpFunc
	// VersionFunc is an instance of a mock function object controlling the
	// behavior of the method Version.
	VersionFunc *StoreVersionFunc
}

// NewMockStore creates a new mock of the Store interface. All methods
// return zero values for all results, unless overwritten.
func NewMockStore() *MockStore {
	return &MockStore{
		DownFunc: &StoreDownFunc{
			defaultHook: func(context.Context, definition.Definition) error {
				return nil
			},
		},
		LockFunc: &StoreLockFunc{
			defaultHook: func(context.Context) (bool, func(err error) error, error) {
				return false, nil, nil
			},
		},
		UpFunc: &StoreUpFunc{
			defaultHook: func(context.Context, definition.Definition) error {
				return nil
			},
		},
		VersionFunc: &StoreVersionFunc{
			defaultHook: func(context.Context) (int, bool, bool, error) {
				return 0, false, false, nil
			},
		},
	}
}

// NewStrictMockStore creates a new mock of the Store interface. All methods
// panic on invocation, unless overwritten.
func NewStrictMockStore() *MockStore {
	return &MockStore{
		DownFunc: &StoreDownFunc{
			defaultHook: func(context.Context, definition.Definition) error {
				panic("unexpected invocation of MockStore.Down")
			},
		},
		LockFunc: &StoreLockFunc{
			defaultHook: func(context.Context) (bool, func(err error) error, error) {
				panic("unexpected invocation of MockStore.Lock")
			},
		},
		UpFunc: &StoreUpFunc{
			defaultHook: func(context.Context, definition.Definition) error {
				panic("unexpected invocation of MockStore.Up")
			},
		},
		VersionFunc: &StoreVersionFunc{
			defaultHook: func(context.Context) (int, bool, bool, error) {
				panic("unexpected invocation of MockStore.Version")
			},
		},
	}
}

// NewMockStoreFrom creates a new mock of the MockStore interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockStoreFrom(i Store) *MockStore {
	return &MockStore{
		DownFunc: &StoreDownFunc{
			defaultHook: i.Down,
		},
		LockFunc: &StoreLockFunc{
			defaultHook: i.Lock,
		},
		UpFunc: &StoreUpFunc{
			defaultHook: i.Up,
		},
		VersionFunc: &StoreVersionFunc{
			defaultHook: i.Version,
		},
	}
}

// StoreDownFunc describes the behavior when the Down method of the parent
// MockStore instance is invoked.
type StoreDownFunc struct {
	defaultHook func(context.Context, definition.Definition) error
	hooks       []func(context.Context, definition.Definition) error
	history     []StoreDownFuncCall
	mutex       sync.Mutex
}

// Down delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) Down(v0 context.Context, v1 definition.Definition) error {
	r0 := m.DownFunc.nextHook()(v0, v1)
	m.DownFunc.appendCall(StoreDownFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Down method of the
// parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreDownFunc) SetDefaultHook(hook func(context.Context, definition.Definition) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Down method of the parent MockStore instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StoreDownFunc) PushHook(hook func(context.Context, definition.Definition) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *StoreDownFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, definition.Definition) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *StoreDownFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, definition.Definition) error {
		return r0
	})
}

func (f *StoreDownFunc) nextHook() func(context.Context, definition.Definition) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreDownFunc) appendCall(r0 StoreDownFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreDownFuncCall objects describing the
// invocations of this function.
func (f *StoreDownFunc) History() []StoreDownFuncCall {
	f.mutex.Lock()
	history := make([]StoreDownFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreDownFuncCall is an object that describes an invocation of method
// Down on an instance of MockStore.
type StoreDownFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 definition.Definition
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreDownFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreDownFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// StoreLockFunc describes the behavior when the Lock method of the parent
// MockStore instance is invoked.
type StoreLockFunc struct {
	defaultHook func(context.Context) (bool, func(err error) error, error)
	hooks       []func(context.Context) (bool, func(err error) error, error)
	history     []StoreLockFuncCall
	mutex       sync.Mutex
}

// Lock delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) Lock(v0 context.Context) (bool, func(err error) error, error) {
	r0, r1, r2 := m.LockFunc.nextHook()(v0)
	m.LockFunc.appendCall(StoreLockFuncCall{v0, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the Lock method of the
// parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreLockFunc) SetDefaultHook(hook func(context.Context) (bool, func(err error) error, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Lock method of the parent MockStore instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StoreLockFunc) PushHook(hook func(context.Context) (bool, func(err error) error, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *StoreLockFunc) SetDefaultReturn(r0 bool, r1 func(err error) error, r2 error) {
	f.SetDefaultHook(func(context.Context) (bool, func(err error) error, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *StoreLockFunc) PushReturn(r0 bool, r1 func(err error) error, r2 error) {
	f.PushHook(func(context.Context) (bool, func(err error) error, error) {
		return r0, r1, r2
	})
}

func (f *StoreLockFunc) nextHook() func(context.Context) (bool, func(err error) error, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreLockFunc) appendCall(r0 StoreLockFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreLockFuncCall objects describing the
// invocations of this function.
func (f *StoreLockFunc) History() []StoreLockFuncCall {
	f.mutex.Lock()
	history := make([]StoreLockFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreLockFuncCall is an object that describes an invocation of method
// Lock on an instance of MockStore.
type StoreLockFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 bool
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 func(err error) error
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreLockFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreLockFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// StoreUpFunc describes the behavior when the Up method of the parent
// MockStore instance is invoked.
type StoreUpFunc struct {
	defaultHook func(context.Context, definition.Definition) error
	hooks       []func(context.Context, definition.Definition) error
	history     []StoreUpFuncCall
	mutex       sync.Mutex
}

// Up delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) Up(v0 context.Context, v1 definition.Definition) error {
	r0 := m.UpFunc.nextHook()(v0, v1)
	m.UpFunc.appendCall(StoreUpFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Up method of the
// parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreUpFunc) SetDefaultHook(hook func(context.Context, definition.Definition) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Up method of the parent MockStore instance invokes the hook at the front
// of the queue and discards it. After the queue is empty, the default hook
// function is invoked for any future action.
func (f *StoreUpFunc) PushHook(hook func(context.Context, definition.Definition) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *StoreUpFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, definition.Definition) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *StoreUpFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, definition.Definition) error {
		return r0
	})
}

func (f *StoreUpFunc) nextHook() func(context.Context, definition.Definition) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreUpFunc) appendCall(r0 StoreUpFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreUpFuncCall objects describing the
// invocations of this function.
func (f *StoreUpFunc) History() []StoreUpFuncCall {
	f.mutex.Lock()
	history := make([]StoreUpFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreUpFuncCall is an object that describes an invocation of method Up on
// an instance of MockStore.
type StoreUpFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 definition.Definition
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreUpFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreUpFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// StoreVersionFunc describes the behavior when the Version method of the
// parent MockStore instance is invoked.
type StoreVersionFunc struct {
	defaultHook func(context.Context) (int, bool, bool, error)
	hooks       []func(context.Context) (int, bool, bool, error)
	history     []StoreVersionFuncCall
	mutex       sync.Mutex
}

// Version delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockStore) Version(v0 context.Context) (int, bool, bool, error) {
	r0, r1, r2, r3 := m.VersionFunc.nextHook()(v0)
	m.VersionFunc.appendCall(StoreVersionFuncCall{v0, r0, r1, r2, r3})
	return r0, r1, r2, r3
}

// SetDefaultHook sets function that is called when the Version method of
// the parent MockStore instance is invoked and the hook queue is empty.
func (f *StoreVersionFunc) SetDefaultHook(hook func(context.Context) (int, bool, bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Version method of the parent MockStore instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *StoreVersionFunc) PushHook(hook func(context.Context) (int, bool, bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *StoreVersionFunc) SetDefaultReturn(r0 int, r1 bool, r2 bool, r3 error) {
	f.SetDefaultHook(func(context.Context) (int, bool, bool, error) {
		return r0, r1, r2, r3
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *StoreVersionFunc) PushReturn(r0 int, r1 bool, r2 bool, r3 error) {
	f.PushHook(func(context.Context) (int, bool, bool, error) {
		return r0, r1, r2, r3
	})
}

func (f *StoreVersionFunc) nextHook() func(context.Context) (int, bool, bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreVersionFunc) appendCall(r0 StoreVersionFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreVersionFuncCall objects describing the
// invocations of this function.
func (f *StoreVersionFunc) History() []StoreVersionFuncCall {
	f.mutex.Lock()
	history := make([]StoreVersionFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreVersionFuncCall is an object that describes an invocation of method
// Version on an instance of MockStore.
type StoreVersionFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 int
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 bool
	// Result3 is the value of the 4th result returned from this method
	// invocation.
	Result3 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreVersionFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreVersionFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2, c.Result3}
}
