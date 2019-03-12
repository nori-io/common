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

import http "net/http"

import mock "github.com/stretchr/testify/mock"
import mux "github.com/gorilla/mux"

// Http is an autogenerated mock type for the Http type
type Http struct {
	mock.Mock
}

// Handle provides a mock function with given fields: pattern, handler
func (_m *Http) Handle(pattern string, handler http.Handler) *mux.Route {
	ret := _m.Called(pattern, handler)

	var r0 *mux.Route
	if rf, ok := ret.Get(0).(func(string, http.Handler) *mux.Route); ok {
		r0 = rf(pattern, handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mux.Route)
		}
	}

	return r0
}

// HandleFunc provides a mock function with given fields: pattern, handler
func (_m *Http) HandleFunc(pattern string, handler http.HandlerFunc) *mux.Route {
	ret := _m.Called(pattern, handler)

	var r0 *mux.Route
	if rf, ok := ret.Get(0).(func(string, http.HandlerFunc) *mux.Route); ok {
		r0 = rf(pattern, handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mux.Route)
		}
	}

	return r0
}
