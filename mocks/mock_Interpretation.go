// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockInterpretation is an autogenerated mock type for the Interpretation type
type MockInterpretation struct {
	mock.Mock
}

type MockInterpretation_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterpretation) EXPECT() *MockInterpretation_Expecter {
	return &MockInterpretation_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: intp
func (_m *MockInterpretation) Add(intp any) {
	_m.Called(intp)
}

// MockInterpretation_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MockInterpretation_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - intp any
func (_e *MockInterpretation_Expecter) Add(intp interface{}) *MockInterpretation_Add_Call {
	return &MockInterpretation_Add_Call{Call: _e.mock.On("Add", intp)}
}

func (_c *MockInterpretation_Add_Call) Run(run func(intp any)) *MockInterpretation_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(any))
	})
	return _c
}

func (_c *MockInterpretation_Add_Call) Return() *MockInterpretation_Add_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockInterpretation_Add_Call) RunAndReturn(run func(any)) *MockInterpretation_Add_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterpretation creates a new instance of MockInterpretation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterpretation(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterpretation {
	mock := &MockInterpretation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}