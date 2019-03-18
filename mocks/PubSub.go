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

import interfaces "github.com/nori-io/nori-common/interfaces"
import mock "github.com/stretchr/testify/mock"

// PubSub is an autogenerated mock type for the PubSub type
type PubSub struct {
	mock.Mock
}

// NewPublisher provides a mock function with given fields: _a0, _a1
func (_m *PubSub) NewPublisher(_a0 interfaces.Marshaller, _a1 ...func(interface{})) (interfaces.Publisher, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 interfaces.Publisher
	if rf, ok := ret.Get(0).(func(interfaces.Marshaller, ...func(interface{})) interfaces.Publisher); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.Publisher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interfaces.Marshaller, ...func(interface{})) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSubscriber provides a mock function with given fields: _a0, _a1
func (_m *PubSub) NewSubscriber(_a0 interfaces.Unmarshaller, _a1 ...func(interface{})) (interfaces.Subscriber, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 interfaces.Subscriber
	if rf, ok := ret.Get(0).(func(interfaces.Unmarshaller, ...func(interface{})) interfaces.Subscriber); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.Subscriber)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interfaces.Unmarshaller, ...func(interface{})) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Publisher provides a mock function with given fields:
func (_m *PubSub) Publisher() interfaces.Publisher {
	ret := _m.Called()

	var r0 interfaces.Publisher
	if rf, ok := ret.Get(0).(func() interfaces.Publisher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.Publisher)
		}
	}

	return r0
}

// Subscriber provides a mock function with given fields:
func (_m *PubSub) Subscriber() interfaces.Subscriber {
	ret := _m.Called()

	var r0 interfaces.Subscriber
	if rf, ok := ret.Get(0).(func() interfaces.Subscriber); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.Subscriber)
		}
	}

	return r0
}