// Code generated by mockery v1.0.0. DO NOT EDIT.

//
/*
Copyright 2019 The Nori Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mocks

import mock "github.com/stretchr/testify/mock"

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// Bool provides a mock function with given fields: key
func (_m *Config) Bool(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Float provides a mock function with given fields: key
func (_m *Config) Float(key string) float64 {
	ret := _m.Called(key)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *Config) Get(key string) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Int provides a mock function with given fields: key
func (_m *Config) Int(key string) int {
	ret := _m.Called(key)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// IsSet provides a mock function with given fields: key
func (_m *Config) IsSet(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetDefault provides a mock function with given fields: key, val
func (_m *Config) SetDefault(key string, val interface{}) {
	_m.Called(key, val)
}

// Slice provides a mock function with given fields: key, delimiter
func (_m *Config) Slice(key string, delimiter string) []interface{} {
	ret := _m.Called(key, delimiter)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(string, string) []interface{}); ok {
		r0 = rf(key, delimiter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// String provides a mock function with given fields: key
func (_m *Config) String(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StringMap provides a mock function with given fields: key
func (_m *Config) StringMap(key string) map[string]interface{} {
	ret := _m.Called(key)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// UInt provides a mock function with given fields: key
func (_m *Config) UInt(key string) uint {
	ret := _m.Called(key)

	var r0 uint
	if rf, ok := ret.Get(0).(func(string) uint); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// Unmarshal provides a mock function with given fields: v, prefix
func (_m *Config) Unmarshal(v interface{}, prefix string) error {
	ret := _m.Called(v, prefix)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string) error); ok {
		r0 = rf(v, prefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}