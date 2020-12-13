// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v4/pkg/domain/plugin (interfaces: Plugin)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "github.com/nori-io/common/v4/pkg/domain/config"
	logger "github.com/nori-io/common/v4/pkg/domain/logger"
	meta "github.com/nori-io/common/v4/pkg/domain/meta"
	registry "github.com/nori-io/common/v4/pkg/domain/registry"
)

// MockPlugin is a mock of Plugin interface
type MockPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockPluginMockRecorder
}

// MockPluginMockRecorder is the mock recorder for MockPlugin
type MockPluginMockRecorder struct {
	mock *MockPlugin
}

// NewMockPlugin creates a new mock instance
func NewMockPlugin(ctrl *gomock.Controller) *MockPlugin {
	mock := &MockPlugin{ctrl: ctrl}
	mock.recorder = &MockPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlugin) EXPECT() *MockPluginMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockPlugin) Init(arg0 context.Context, arg1 config.Config, arg2 logger.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockPluginMockRecorder) Init(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockPlugin)(nil).Init), arg0, arg1, arg2)
}

// Instance mocks base method
func (m *MockPlugin) Instance() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Instance")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Instance indicates an expected call of Instance
func (mr *MockPluginMockRecorder) Instance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Instance", reflect.TypeOf((*MockPlugin)(nil).Instance))
}

// Meta mocks base method
func (m *MockPlugin) Meta() meta.Meta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(meta.Meta)
	return ret0
}

// Meta indicates an expected call of Meta
func (mr *MockPluginMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockPlugin)(nil).Meta))
}

// Start mocks base method
func (m *MockPlugin) Start(arg0 context.Context, arg1 registry.Registry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockPluginMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockPlugin)(nil).Start), arg0, arg1)
}

// Stop mocks base method
func (m *MockPlugin) Stop(arg0 context.Context, arg1 registry.Registry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockPluginMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockPlugin)(nil).Stop), arg0, arg1)
}
