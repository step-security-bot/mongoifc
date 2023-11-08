// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	bson "go.mongodb.org/mongo-driver/bson"

	mock "github.com/stretchr/testify/mock"

	mongoifc "github.com/sv-tools/mongoifc"

	options "go.mongodb.org/mongo-driver/mongo/options"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	time "time"
)

// SessionContext is an autogenerated mock type for the SessionContext type
type SessionContext struct {
	mock.Mock
}

type SessionContext_Expecter struct {
	mock *mock.Mock
}

func (_m *SessionContext) EXPECT() *SessionContext_Expecter {
	return &SessionContext_Expecter{mock: &_m.Mock}
}

// AbortTransaction provides a mock function with given fields: ctx
func (_m *SessionContext) AbortTransaction(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionContext_AbortTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AbortTransaction'
type SessionContext_AbortTransaction_Call struct {
	*mock.Call
}

// AbortTransaction is a helper method to define mock.On call
//   - ctx context.Context
func (_e *SessionContext_Expecter) AbortTransaction(ctx interface{}) *SessionContext_AbortTransaction_Call {
	return &SessionContext_AbortTransaction_Call{Call: _e.mock.On("AbortTransaction", ctx)}
}

func (_c *SessionContext_AbortTransaction_Call) Run(run func(ctx context.Context)) *SessionContext_AbortTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SessionContext_AbortTransaction_Call) Return(_a0 error) *SessionContext_AbortTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_AbortTransaction_Call) RunAndReturn(run func(context.Context) error) *SessionContext_AbortTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// AdvanceClusterTime provides a mock function with given fields: _a0
func (_m *SessionContext) AdvanceClusterTime(_a0 bson.Raw) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(bson.Raw) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionContext_AdvanceClusterTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AdvanceClusterTime'
type SessionContext_AdvanceClusterTime_Call struct {
	*mock.Call
}

// AdvanceClusterTime is a helper method to define mock.On call
//   - _a0 bson.Raw
func (_e *SessionContext_Expecter) AdvanceClusterTime(_a0 interface{}) *SessionContext_AdvanceClusterTime_Call {
	return &SessionContext_AdvanceClusterTime_Call{Call: _e.mock.On("AdvanceClusterTime", _a0)}
}

func (_c *SessionContext_AdvanceClusterTime_Call) Run(run func(_a0 bson.Raw)) *SessionContext_AdvanceClusterTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bson.Raw))
	})
	return _c
}

func (_c *SessionContext_AdvanceClusterTime_Call) Return(_a0 error) *SessionContext_AdvanceClusterTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_AdvanceClusterTime_Call) RunAndReturn(run func(bson.Raw) error) *SessionContext_AdvanceClusterTime_Call {
	_c.Call.Return(run)
	return _c
}

// AdvanceOperationTime provides a mock function with given fields: _a0
func (_m *SessionContext) AdvanceOperationTime(_a0 *primitive.Timestamp) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*primitive.Timestamp) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionContext_AdvanceOperationTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AdvanceOperationTime'
type SessionContext_AdvanceOperationTime_Call struct {
	*mock.Call
}

// AdvanceOperationTime is a helper method to define mock.On call
//   - _a0 *primitive.Timestamp
func (_e *SessionContext_Expecter) AdvanceOperationTime(_a0 interface{}) *SessionContext_AdvanceOperationTime_Call {
	return &SessionContext_AdvanceOperationTime_Call{Call: _e.mock.On("AdvanceOperationTime", _a0)}
}

func (_c *SessionContext_AdvanceOperationTime_Call) Run(run func(_a0 *primitive.Timestamp)) *SessionContext_AdvanceOperationTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*primitive.Timestamp))
	})
	return _c
}

func (_c *SessionContext_AdvanceOperationTime_Call) Return(_a0 error) *SessionContext_AdvanceOperationTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_AdvanceOperationTime_Call) RunAndReturn(run func(*primitive.Timestamp) error) *SessionContext_AdvanceOperationTime_Call {
	_c.Call.Return(run)
	return _c
}

// Client provides a mock function with given fields:
func (_m *SessionContext) Client() mongoifc.Client {
	ret := _m.Called()

	var r0 mongoifc.Client
	if rf, ok := ret.Get(0).(func() mongoifc.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.Client)
		}
	}

	return r0
}

// SessionContext_Client_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Client'
type SessionContext_Client_Call struct {
	*mock.Call
}

// Client is a helper method to define mock.On call
func (_e *SessionContext_Expecter) Client() *SessionContext_Client_Call {
	return &SessionContext_Client_Call{Call: _e.mock.On("Client")}
}

func (_c *SessionContext_Client_Call) Run(run func()) *SessionContext_Client_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SessionContext_Client_Call) Return(_a0 mongoifc.Client) *SessionContext_Client_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_Client_Call) RunAndReturn(run func() mongoifc.Client) *SessionContext_Client_Call {
	_c.Call.Return(run)
	return _c
}

// ClusterTime provides a mock function with given fields:
func (_m *SessionContext) ClusterTime() bson.Raw {
	ret := _m.Called()

	var r0 bson.Raw
	if rf, ok := ret.Get(0).(func() bson.Raw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bson.Raw)
		}
	}

	return r0
}

// SessionContext_ClusterTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClusterTime'
type SessionContext_ClusterTime_Call struct {
	*mock.Call
}

// ClusterTime is a helper method to define mock.On call
func (_e *SessionContext_Expecter) ClusterTime() *SessionContext_ClusterTime_Call {
	return &SessionContext_ClusterTime_Call{Call: _e.mock.On("ClusterTime")}
}

func (_c *SessionContext_ClusterTime_Call) Run(run func()) *SessionContext_ClusterTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SessionContext_ClusterTime_Call) Return(_a0 bson.Raw) *SessionContext_ClusterTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_ClusterTime_Call) RunAndReturn(run func() bson.Raw) *SessionContext_ClusterTime_Call {
	_c.Call.Return(run)
	return _c
}

// CommitTransaction provides a mock function with given fields: ctx
func (_m *SessionContext) CommitTransaction(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionContext_CommitTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CommitTransaction'
type SessionContext_CommitTransaction_Call struct {
	*mock.Call
}

// CommitTransaction is a helper method to define mock.On call
//   - ctx context.Context
func (_e *SessionContext_Expecter) CommitTransaction(ctx interface{}) *SessionContext_CommitTransaction_Call {
	return &SessionContext_CommitTransaction_Call{Call: _e.mock.On("CommitTransaction", ctx)}
}

func (_c *SessionContext_CommitTransaction_Call) Run(run func(ctx context.Context)) *SessionContext_CommitTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SessionContext_CommitTransaction_Call) Return(_a0 error) *SessionContext_CommitTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_CommitTransaction_Call) RunAndReturn(run func(context.Context) error) *SessionContext_CommitTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// Deadline provides a mock function with given fields:
func (_m *SessionContext) Deadline() (time.Time, bool) {
	ret := _m.Called()

	var r0 time.Time
	var r1 bool
	if rf, ok := ret.Get(0).(func() (time.Time, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// SessionContext_Deadline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Deadline'
type SessionContext_Deadline_Call struct {
	*mock.Call
}

// Deadline is a helper method to define mock.On call
func (_e *SessionContext_Expecter) Deadline() *SessionContext_Deadline_Call {
	return &SessionContext_Deadline_Call{Call: _e.mock.On("Deadline")}
}

func (_c *SessionContext_Deadline_Call) Run(run func()) *SessionContext_Deadline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SessionContext_Deadline_Call) Return(deadline time.Time, ok bool) *SessionContext_Deadline_Call {
	_c.Call.Return(deadline, ok)
	return _c
}

func (_c *SessionContext_Deadline_Call) RunAndReturn(run func() (time.Time, bool)) *SessionContext_Deadline_Call {
	_c.Call.Return(run)
	return _c
}

// Done provides a mock function with given fields:
func (_m *SessionContext) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// SessionContext_Done_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Done'
type SessionContext_Done_Call struct {
	*mock.Call
}

// Done is a helper method to define mock.On call
func (_e *SessionContext_Expecter) Done() *SessionContext_Done_Call {
	return &SessionContext_Done_Call{Call: _e.mock.On("Done")}
}

func (_c *SessionContext_Done_Call) Run(run func()) *SessionContext_Done_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SessionContext_Done_Call) Return(_a0 <-chan struct{}) *SessionContext_Done_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_Done_Call) RunAndReturn(run func() <-chan struct{}) *SessionContext_Done_Call {
	_c.Call.Return(run)
	return _c
}

// EndSession provides a mock function with given fields: ctx
func (_m *SessionContext) EndSession(ctx context.Context) {
	_m.Called(ctx)
}

// SessionContext_EndSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EndSession'
type SessionContext_EndSession_Call struct {
	*mock.Call
}

// EndSession is a helper method to define mock.On call
//   - ctx context.Context
func (_e *SessionContext_Expecter) EndSession(ctx interface{}) *SessionContext_EndSession_Call {
	return &SessionContext_EndSession_Call{Call: _e.mock.On("EndSession", ctx)}
}

func (_c *SessionContext_EndSession_Call) Run(run func(ctx context.Context)) *SessionContext_EndSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SessionContext_EndSession_Call) Return() *SessionContext_EndSession_Call {
	_c.Call.Return()
	return _c
}

func (_c *SessionContext_EndSession_Call) RunAndReturn(run func(context.Context)) *SessionContext_EndSession_Call {
	_c.Call.Return(run)
	return _c
}

// Err provides a mock function with given fields:
func (_m *SessionContext) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionContext_Err_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Err'
type SessionContext_Err_Call struct {
	*mock.Call
}

// Err is a helper method to define mock.On call
func (_e *SessionContext_Expecter) Err() *SessionContext_Err_Call {
	return &SessionContext_Err_Call{Call: _e.mock.On("Err")}
}

func (_c *SessionContext_Err_Call) Run(run func()) *SessionContext_Err_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SessionContext_Err_Call) Return(_a0 error) *SessionContext_Err_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_Err_Call) RunAndReturn(run func() error) *SessionContext_Err_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with given fields:
func (_m *SessionContext) ID() bson.Raw {
	ret := _m.Called()

	var r0 bson.Raw
	if rf, ok := ret.Get(0).(func() bson.Raw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bson.Raw)
		}
	}

	return r0
}

// SessionContext_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type SessionContext_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *SessionContext_Expecter) ID() *SessionContext_ID_Call {
	return &SessionContext_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *SessionContext_ID_Call) Run(run func()) *SessionContext_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SessionContext_ID_Call) Return(_a0 bson.Raw) *SessionContext_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_ID_Call) RunAndReturn(run func() bson.Raw) *SessionContext_ID_Call {
	_c.Call.Return(run)
	return _c
}

// OperationTime provides a mock function with given fields:
func (_m *SessionContext) OperationTime() *primitive.Timestamp {
	ret := _m.Called()

	var r0 *primitive.Timestamp
	if rf, ok := ret.Get(0).(func() *primitive.Timestamp); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*primitive.Timestamp)
		}
	}

	return r0
}

// SessionContext_OperationTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OperationTime'
type SessionContext_OperationTime_Call struct {
	*mock.Call
}

// OperationTime is a helper method to define mock.On call
func (_e *SessionContext_Expecter) OperationTime() *SessionContext_OperationTime_Call {
	return &SessionContext_OperationTime_Call{Call: _e.mock.On("OperationTime")}
}

func (_c *SessionContext_OperationTime_Call) Run(run func()) *SessionContext_OperationTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SessionContext_OperationTime_Call) Return(_a0 *primitive.Timestamp) *SessionContext_OperationTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_OperationTime_Call) RunAndReturn(run func() *primitive.Timestamp) *SessionContext_OperationTime_Call {
	_c.Call.Return(run)
	return _c
}

// StartTransaction provides a mock function with given fields: opts
func (_m *SessionContext) StartTransaction(opts ...*options.TransactionOptions) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*options.TransactionOptions) error); ok {
		r0 = rf(opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SessionContext_StartTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartTransaction'
type SessionContext_StartTransaction_Call struct {
	*mock.Call
}

// StartTransaction is a helper method to define mock.On call
//   - opts ...*options.TransactionOptions
func (_e *SessionContext_Expecter) StartTransaction(opts ...interface{}) *SessionContext_StartTransaction_Call {
	return &SessionContext_StartTransaction_Call{Call: _e.mock.On("StartTransaction",
		append([]interface{}{}, opts...)...)}
}

func (_c *SessionContext_StartTransaction_Call) Run(run func(opts ...*options.TransactionOptions)) *SessionContext_StartTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.TransactionOptions, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(*options.TransactionOptions)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *SessionContext_StartTransaction_Call) Return(_a0 error) *SessionContext_StartTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_StartTransaction_Call) RunAndReturn(run func(...*options.TransactionOptions) error) *SessionContext_StartTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// Value provides a mock function with given fields: key
func (_m *SessionContext) Value(key interface{}) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// SessionContext_Value_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Value'
type SessionContext_Value_Call struct {
	*mock.Call
}

// Value is a helper method to define mock.On call
//   - key interface{}
func (_e *SessionContext_Expecter) Value(key interface{}) *SessionContext_Value_Call {
	return &SessionContext_Value_Call{Call: _e.mock.On("Value", key)}
}

func (_c *SessionContext_Value_Call) Run(run func(key interface{})) *SessionContext_Value_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *SessionContext_Value_Call) Return(_a0 interface{}) *SessionContext_Value_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SessionContext_Value_Call) RunAndReturn(run func(interface{}) interface{}) *SessionContext_Value_Call {
	_c.Call.Return(run)
	return _c
}

// WithTransaction provides a mock function with given fields: ctx, fn, opts
func (_m *SessionContext) WithTransaction(ctx context.Context, fn func(mongoifc.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, fn)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, func(mongoifc.SessionContext) (interface{}, error), ...*options.TransactionOptions) (interface{}, error)); ok {
		return rf(ctx, fn, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, func(mongoifc.SessionContext) (interface{}, error), ...*options.TransactionOptions) interface{}); ok {
		r0 = rf(ctx, fn, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, func(mongoifc.SessionContext) (interface{}, error), ...*options.TransactionOptions) error); ok {
		r1 = rf(ctx, fn, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SessionContext_WithTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTransaction'
type SessionContext_WithTransaction_Call struct {
	*mock.Call
}

// WithTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - fn func(mongoifc.SessionContext)(interface{} , error)
//   - opts ...*options.TransactionOptions
func (_e *SessionContext_Expecter) WithTransaction(ctx interface{}, fn interface{}, opts ...interface{}) *SessionContext_WithTransaction_Call {
	return &SessionContext_WithTransaction_Call{Call: _e.mock.On("WithTransaction",
		append([]interface{}{ctx, fn}, opts...)...)}
}

func (_c *SessionContext_WithTransaction_Call) Run(run func(ctx context.Context, fn func(mongoifc.SessionContext) (interface{}, error), opts ...*options.TransactionOptions)) *SessionContext_WithTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.TransactionOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*options.TransactionOptions)
			}
		}
		run(args[0].(context.Context), args[1].(func(mongoifc.SessionContext) (interface{}, error)), variadicArgs...)
	})
	return _c
}

func (_c *SessionContext_WithTransaction_Call) Return(_a0 interface{}, _a1 error) *SessionContext_WithTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SessionContext_WithTransaction_Call) RunAndReturn(run func(context.Context, func(mongoifc.SessionContext) (interface{}, error), ...*options.TransactionOptions) (interface{}, error)) *SessionContext_WithTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewSessionContext creates a new instance of SessionContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSessionContext(t interface {
	mock.TestingT
	Cleanup(func())
}) *SessionContext {
	mock := &SessionContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
