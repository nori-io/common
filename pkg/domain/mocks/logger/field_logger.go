// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v4/pkg/domain/logger (interfaces: FieldLogger)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logger "github.com/nori-io/common/v4/pkg/domain/logger"
)

// MockFieldLogger is a mock of FieldLogger interface
type MockFieldLogger struct {
	ctrl     *gomock.Controller
	recorder *MockFieldLoggerMockRecorder
}

// MockFieldLoggerMockRecorder is the mock recorder for MockFieldLogger
type MockFieldLoggerMockRecorder struct {
	mock *MockFieldLogger
}

// NewMockFieldLogger creates a new mock instance
func NewMockFieldLogger(ctrl *gomock.Controller) *MockFieldLogger {
	mock := &MockFieldLogger{ctrl: ctrl}
	mock.recorder = &MockFieldLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFieldLogger) EXPECT() *MockFieldLoggerMockRecorder {
	return m.recorder
}

// Critical mocks base method
func (m *MockFieldLogger) Critical(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Critical", varargs...)
}

// Critical indicates an expected call of Critical
func (mr *MockFieldLoggerMockRecorder) Critical(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Critical", reflect.TypeOf((*MockFieldLogger)(nil).Critical), varargs...)
}

// Debug mocks base method
func (m *MockFieldLogger) Debug(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockFieldLoggerMockRecorder) Debug(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockFieldLogger)(nil).Debug), varargs...)
}

// Error mocks base method
func (m *MockFieldLogger) Error(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockFieldLoggerMockRecorder) Error(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockFieldLogger)(nil).Error), varargs...)
}

// Fatal mocks base method
func (m *MockFieldLogger) Fatal(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal
func (mr *MockFieldLoggerMockRecorder) Fatal(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockFieldLogger)(nil).Fatal), varargs...)
}

// Info mocks base method
func (m *MockFieldLogger) Info(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockFieldLoggerMockRecorder) Info(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockFieldLogger)(nil).Info), varargs...)
}

// Log mocks base method
func (m *MockFieldLogger) Log(arg0 logger.Level, arg1 string, arg2 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log
func (mr *MockFieldLoggerMockRecorder) Log(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockFieldLogger)(nil).Log), varargs...)
}

// Notice mocks base method
func (m *MockFieldLogger) Notice(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Notice", varargs...)
}

// Notice indicates an expected call of Notice
func (mr *MockFieldLoggerMockRecorder) Notice(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notice", reflect.TypeOf((*MockFieldLogger)(nil).Notice), varargs...)
}

// Panic mocks base method
func (m *MockFieldLogger) Panic(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Panic", varargs...)
}

// Panic indicates an expected call of Panic
func (mr *MockFieldLoggerMockRecorder) Panic(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panic", reflect.TypeOf((*MockFieldLogger)(nil).Panic), varargs...)
}

// Warning mocks base method
func (m *MockFieldLogger) Warning(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warning", varargs...)
}

// Warning indicates an expected call of Warning
func (mr *MockFieldLoggerMockRecorder) Warning(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockFieldLogger)(nil).Warning), varargs...)
}
