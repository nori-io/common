// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v4/pkg/domain/registry (interfaces: Registry)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	meta "github.com/nori-io/common/v4/pkg/domain/meta"
)

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockRegistry) ID(arg0 meta.ID) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID
func (mr *MockRegistryMockRecorder) ID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockRegistry)(nil).ID), arg0)
}

// Interface mocks base method
func (m *MockRegistry) Interface(arg0 meta.Interface) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Interface", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Interface indicates an expected call of Interface
func (mr *MockRegistryMockRecorder) Interface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Interface", reflect.TypeOf((*MockRegistry)(nil).Interface), arg0)
}

// Resolve mocks base method
func (m *MockRegistry) Resolve(arg0 meta.Dependency) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (mr *MockRegistryMockRecorder) Resolve(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockRegistry)(nil).Resolve), arg0)
}
