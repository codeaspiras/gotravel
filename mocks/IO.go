// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IO is an autogenerated mock type for the IO type
type IO struct {
	mock.Mock
}

// Ask provides a mock function with given fields: template, arguments
func (_m *IO) Ask(template string, arguments ...interface{}) (string, error) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, arguments...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Ask")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) (string, error)); ok {
		return rf(template, arguments...)
	}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = rf(template, arguments...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(template, arguments...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Echo provides a mock function with given fields: template, arguments
func (_m *IO) Echo(template string, arguments ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, arguments...)
	_m.Called(_ca...)
}

// NewIO creates a new instance of IO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIO(t interface {
	mock.TestingT
	Cleanup(func())
}) *IO {
	mock := &IO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}