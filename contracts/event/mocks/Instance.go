// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	event "github.com/mewway/go-laravel/contracts/event"
	mock "github.com/stretchr/testify/mock"
)

// Instance is an autogenerated mock type for the Instance type
type Instance struct {
	mock.Mock
}

// GetEvents provides a mock function with given fields:
func (_m *Instance) GetEvents() map[event.Event][]event.Listener {
	ret := _m.Called()

	var r0 map[event.Event][]event.Listener
	if rf, ok := ret.Get(0).(func() map[event.Event][]event.Listener); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[event.Event][]event.Listener)
		}
	}

	return r0
}

// Job provides a mock function with given fields: _a0, args
func (_m *Instance) Job(_a0 event.Event, args []event.Arg) event.Task {
	ret := _m.Called(_a0, args)

	var r0 event.Task
	if rf, ok := ret.Get(0).(func(event.Event, []event.Arg) event.Task); ok {
		r0 = rf(_a0, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Task)
		}
	}

	return r0
}

// Register provides a mock function with given fields: _a0
func (_m *Instance) Register(_a0 map[event.Event][]event.Listener) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewInstance interface {
	mock.TestingT
	Cleanup(func())
}

// NewInstance creates a new instance of Instance. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInstance(t mockConstructorTestingTNewInstance) *Instance {
	mock := &Instance{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
