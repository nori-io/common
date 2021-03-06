// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v5/pkg/domain/registry (interfaces: ConfigRegistry)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "github.com/nori-io/common/v5/pkg/domain/config"
	meta "github.com/nori-io/common/v5/pkg/domain/meta"
	registry "github.com/nori-io/common/v5/pkg/domain/registry"
)

// MockConfigRegistry is a mock of ConfigRegistry interface.
type MockConfigRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockConfigRegistryMockRecorder
}

// MockConfigRegistryMockRecorder is the mock recorder for MockConfigRegistry.
type MockConfigRegistryMockRecorder struct {
	mock *MockConfigRegistry
}

// NewMockConfigRegistry creates a new mock instance.
func NewMockConfigRegistry(ctrl *gomock.Controller) *MockConfigRegistry {
	mock := &MockConfigRegistry{ctrl: ctrl}
	mock.recorder = &MockConfigRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigRegistry) EXPECT() *MockConfigRegistryMockRecorder {
	return m.recorder
}

// PluginVariables mocks base method.
func (m *MockConfigRegistry) PluginVariables(arg0 meta.ID) []registry.Variable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PluginVariables", arg0)
	ret0, _ := ret[0].([]registry.Variable)
	return ret0
}

// PluginVariables indicates an expected call of PluginVariables.
func (mr *MockConfigRegistryMockRecorder) PluginVariables(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PluginVariables", reflect.TypeOf((*MockConfigRegistry)(nil).PluginVariables), arg0)
}

// Register mocks base method.
func (m *MockConfigRegistry) Register(arg0 meta.ID) config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(config.Config)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockConfigRegistryMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockConfigRegistry)(nil).Register), arg0)
}
