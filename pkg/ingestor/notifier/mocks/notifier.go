// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Notifier is an autogenerated mock type for the Notifier type
type Notifier struct {
	mock.Mock
}

type Notifier_Expecter struct {
	mock *mock.Mock
}

func (_m *Notifier) EXPECT() *Notifier_Expecter {
	return &Notifier_Expecter{mock: &_m.Mock}
}

// Notify provides a mock function with given fields: ctx, clusterName, runID
func (_m *Notifier) Notify(ctx context.Context, clusterName string, runID string) error {
	ret := _m.Called(ctx, clusterName, runID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, clusterName, runID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Notifier_Notify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Notify'
type Notifier_Notify_Call struct {
	*mock.Call
}

// Notify is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - runID string
func (_e *Notifier_Expecter) Notify(ctx interface{}, clusterName interface{}, runID interface{}) *Notifier_Notify_Call {
	return &Notifier_Notify_Call{Call: _e.mock.On("Notify", ctx, clusterName, runID)}
}

func (_c *Notifier_Notify_Call) Run(run func(ctx context.Context, clusterName string, runID string)) *Notifier_Notify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Notifier_Notify_Call) Return(_a0 error) *Notifier_Notify_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Notifier_Notify_Call) RunAndReturn(run func(context.Context, string, string) error) *Notifier_Notify_Call {
	_c.Call.Return(run)
	return _c
}

// NewNotifier creates a new instance of Notifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotifier(t interface {
	mock.TestingT
	Cleanup(func())
}) *Notifier {
	mock := &Notifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
