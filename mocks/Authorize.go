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

import context "context"

import mock "github.com/stretchr/testify/mock"

// Authorize is an autogenerated mock type for the Authorize type
type Authorize struct {
	mock.Mock
}

// Authorize provides a mock function with given fields: _a0, _a1
func (_m *Authorize) Authorize(_a0 context.Context, _a1 ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) error); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}