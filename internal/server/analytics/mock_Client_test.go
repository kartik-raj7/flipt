// Code generated by mockery v2.47.0. DO NOT EDIT.

package analytics

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// GetFlagEvaluationsCount provides a mock function with given fields: ctx, req
func (_m *MockClient) GetFlagEvaluationsCount(ctx context.Context, req *FlagEvaluationsCountRequest) ([]string, []float32, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetFlagEvaluationsCount")
	}

	var r0 []string
	var r1 []float32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *FlagEvaluationsCountRequest) ([]string, []float32, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *FlagEvaluationsCountRequest) []string); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *FlagEvaluationsCountRequest) []float32); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]float32)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *FlagEvaluationsCountRequest) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockClient_GetFlagEvaluationsCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlagEvaluationsCount'
type MockClient_GetFlagEvaluationsCount_Call struct {
	*mock.Call
}

// GetFlagEvaluationsCount is a helper method to define mock.On call
//   - ctx context.Context
//   - req *FlagEvaluationsCountRequest
func (_e *MockClient_Expecter) GetFlagEvaluationsCount(ctx interface{}, req interface{}) *MockClient_GetFlagEvaluationsCount_Call {
	return &MockClient_GetFlagEvaluationsCount_Call{Call: _e.mock.On("GetFlagEvaluationsCount", ctx, req)}
}

func (_c *MockClient_GetFlagEvaluationsCount_Call) Run(run func(ctx context.Context, req *FlagEvaluationsCountRequest)) *MockClient_GetFlagEvaluationsCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*FlagEvaluationsCountRequest))
	})
	return _c
}

func (_c *MockClient_GetFlagEvaluationsCount_Call) Return(_a0 []string, _a1 []float32, _a2 error) *MockClient_GetFlagEvaluationsCount_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockClient_GetFlagEvaluationsCount_Call) RunAndReturn(run func(context.Context, *FlagEvaluationsCountRequest) ([]string, []float32, error)) *MockClient_GetFlagEvaluationsCount_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockClient) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockClient_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockClient_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockClient_Expecter) String() *MockClient_String_Call {
	return &MockClient_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockClient_String_Call) Run(run func()) *MockClient_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClient_String_Call) Return(_a0 string) *MockClient_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_String_Call) RunAndReturn(run func() string) *MockClient_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
