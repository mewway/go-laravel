// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	queue "github.com/mewway/go-laravel/contracts/queue"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Task is an autogenerated mock type for the Task type
type Task struct {
	mock.Mock
}

// Delay provides a mock function with given fields: _a0
func (_m *Task) Delay(_a0 time.Time) queue.Task {
	ret := _m.Called(_a0)

	var r0 queue.Task
	if rf, ok := ret.Get(0).(func(time.Time) queue.Task); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queue.Task)
		}
	}

	return r0
}

// Dispatch provides a mock function with given fields:
func (_m *Task) Dispatch() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DispatchSync provides a mock function with given fields:
func (_m *Task) DispatchSync() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnConnection provides a mock function with given fields: connection
func (_m *Task) OnConnection(connection string) queue.Task {
	ret := _m.Called(connection)

	var r0 queue.Task
	if rf, ok := ret.Get(0).(func(string) queue.Task); ok {
		r0 = rf(connection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queue.Task)
		}
	}

	return r0
}

// OnQueue provides a mock function with given fields: _a0
func (_m *Task) OnQueue(_a0 string) queue.Task {
	ret := _m.Called(_a0)

	var r0 queue.Task
	if rf, ok := ret.Get(0).(func(string) queue.Task); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queue.Task)
		}
	}

	return r0
}

type mockConstructorTestingTNewTask interface {
	mock.TestingT
	Cleanup(func())
}

// NewTask creates a new instance of Task. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTask(t mockConstructorTestingTNewTask) *Task {
	mock := &Task{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
