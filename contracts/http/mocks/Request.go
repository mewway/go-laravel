// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	filesystem "github.com/mewway/go-laravel/contracts/filesystem"
	http "github.com/mewway/go-laravel/contracts/http"

	mock "github.com/stretchr/testify/mock"

	nethttp "net/http"

	validation "github.com/mewway/go-laravel/contracts/validation"
)

// Request is an autogenerated mock type for the Request type
type Request struct {
	mock.Mock
}

// AbortWithStatus provides a mock function with given fields: code
func (_m *Request) AbortWithStatus(code int) {
	_m.Called(code)
}

// AbortWithStatusJson provides a mock function with given fields: code, jsonObj
func (_m *Request) AbortWithStatusJson(code int, jsonObj interface{}) {
	_m.Called(code, jsonObj)
}

// Bind provides a mock function with given fields: obj
func (_m *Request) Bind(obj interface{}) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// File provides a mock function with given fields: name
func (_m *Request) File(name string) (filesystem.File, error) {
	ret := _m.Called(name)

	var r0 filesystem.File
	if rf, ok := ret.Get(0).(func(string) filesystem.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filesystem.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Form provides a mock function with given fields: key, defaultValue
func (_m *Request) Form(key string, defaultValue ...string) string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// FullUrl provides a mock function with given fields:
func (_m *Request) FullUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Header provides a mock function with given fields: key, defaultValue
func (_m *Request) Header(key string, defaultValue ...string) string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Headers provides a mock function with given fields:
func (_m *Request) Headers() nethttp.Header {
	ret := _m.Called()

	var r0 nethttp.Header
	if rf, ok := ret.Get(0).(func() nethttp.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nethttp.Header)
		}
	}

	return r0
}

// Input provides a mock function with given fields: key, defaultValue
func (_m *Request) Input(key string, defaultValue ...string) string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// InputBool provides a mock function with given fields: key, defaultValue
func (_m *Request) InputBool(key string, defaultValue ...bool) bool {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...bool) bool); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InputInt provides a mock function with given fields: key, defaultValue
func (_m *Request) InputInt(key string, defaultValue ...int) int {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// InputInt64 provides a mock function with given fields: key, defaultValue
func (_m *Request) InputInt64(key string, defaultValue ...int64) int64 {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, ...int64) int64); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Ip provides a mock function with given fields:
func (_m *Request) Ip() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Json provides a mock function with given fields: key, defaultValue
func (_m *Request) Json(key string, defaultValue ...string) string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Method provides a mock function with given fields:
func (_m *Request) Method() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *Request) Next() {
	_m.Called()
}

// Origin provides a mock function with given fields:
func (_m *Request) Origin() *nethttp.Request {
	ret := _m.Called()

	var r0 *nethttp.Request
	if rf, ok := ret.Get(0).(func() *nethttp.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nethttp.Request)
		}
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *Request) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Query provides a mock function with given fields: key, defaultValue
func (_m *Request) Query(key string, defaultValue ...string) string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// QueryArray provides a mock function with given fields: key
func (_m *Request) QueryArray(key string) []string {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// QueryBool provides a mock function with given fields: key, defaultValue
func (_m *Request) QueryBool(key string, defaultValue ...bool) bool {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...bool) bool); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// QueryInt provides a mock function with given fields: key, defaultValue
func (_m *Request) QueryInt(key string, defaultValue ...int) int {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// QueryInt64 provides a mock function with given fields: key, defaultValue
func (_m *Request) QueryInt64(key string, defaultValue ...int64) int64 {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, ...int64) int64); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// QueryMap provides a mock function with given fields: key
func (_m *Request) QueryMap(key string) map[string]string {
	ret := _m.Called(key)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// Route provides a mock function with given fields: key
func (_m *Request) Route(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RouteInt provides a mock function with given fields: key
func (_m *Request) RouteInt(key string) int {
	ret := _m.Called(key)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// RouteInt64 provides a mock function with given fields: key
func (_m *Request) RouteInt64(key string) int64 {
	ret := _m.Called(key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Url provides a mock function with given fields:
func (_m *Request) Url() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Validate provides a mock function with given fields: rules, options
func (_m *Request) Validate(rules map[string]string, options ...validation.Option) (validation.Validator, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, rules)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 validation.Validator
	if rf, ok := ret.Get(0).(func(map[string]string, ...validation.Option) validation.Validator); ok {
		r0 = rf(rules, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Validator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string, ...validation.Option) error); ok {
		r1 = rf(rules, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateRequest provides a mock function with given fields: request
func (_m *Request) ValidateRequest(request http.FormRequest) (validation.Errors, error) {
	ret := _m.Called(request)

	var r0 validation.Errors
	if rf, ok := ret.Get(0).(func(http.FormRequest) validation.Errors); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Errors)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(http.FormRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRequest interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequest creates a new instance of Request. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequest(t mockConstructorTestingTNewRequest) *Request {
	mock := &Request{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
