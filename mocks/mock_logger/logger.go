// Code generated by MockGen. DO NOT EDIT.
// Source: logger/logger.go

// Package mock_logger is a generated GoMock package.
package mock_logger

import (
	gomock "github.com/golang/mock/gomock"
	logger "github.com/nori-io/nori-common/logger"
	reflect "reflect"
)

// MockHook is a mock of Hook interface
type MockHook struct {
	ctrl     *gomock.Controller
	recorder *MockHookMockRecorder
}

// MockHookMockRecorder is the mock recorder for MockHook
type MockHookMockRecorder struct {
	mock *MockHook
}

// NewMockHook creates a new mock instance
func NewMockHook(ctrl *gomock.Controller) *MockHook {
	mock := &MockHook{ctrl: ctrl}
	mock.recorder = &MockHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHook) EXPECT() *MockHookMockRecorder {
	return m.recorder
}

// Levels mocks base method
func (m *MockHook) Levels() []logger.Level {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Levels")
	ret0, _ := ret[0].([]logger.Level)
	return ret0
}

// Levels indicates an expected call of Levels
func (mr *MockHookMockRecorder) Levels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Levels", reflect.TypeOf((*MockHook)(nil).Levels))
}

// Fire mocks base method
func (m *MockHook) Fire(level logger.Level, message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fire", level, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fire indicates an expected call of Fire
func (mr *MockHookMockRecorder) Fire(level, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fire", reflect.TypeOf((*MockHook)(nil).Fire), level, message)
}

// MockFormatter is a mock of Formatter interface
type MockFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockFormatterMockRecorder
}

// MockFormatterMockRecorder is the mock recorder for MockFormatter
type MockFormatterMockRecorder struct {
	mock *MockFormatter
}

// NewMockFormatter creates a new mock instance
func NewMockFormatter(ctrl *gomock.Controller) *MockFormatter {
	mock := &MockFormatter{ctrl: ctrl}
	mock.recorder = &MockFormatterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFormatter) EXPECT() *MockFormatterMockRecorder {
	return m.recorder
}

// Format mocks base method
func (m *MockFormatter) Format(e logger.Entry, field ...logger.Field) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{e}
	for _, a := range field {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Format", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Format indicates an expected call of Format
func (mr *MockFormatterMockRecorder) Format(e interface{}, field ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{e}, field...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Format", reflect.TypeOf((*MockFormatter)(nil).Format), varargs...)
}

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Critical mocks base method
func (m *MockLogger) Critical(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Critical", varargs...)
}

// Critical indicates an expected call of Critical
func (mr *MockLoggerMockRecorder) Critical(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Critical", reflect.TypeOf((*MockLogger)(nil).Critical), varargs...)
}

// Debug mocks base method
func (m *MockLogger) Debug(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockLoggerMockRecorder) Debug(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), varargs...)
}

// Fatal mocks base method
func (m *MockLogger) Fatal(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal
func (mr *MockLoggerMockRecorder) Fatal(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), varargs...)
}

// Error mocks base method
func (m *MockLogger) Error(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockLoggerMockRecorder) Error(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), varargs...)
}

// Info mocks base method
func (m *MockLogger) Info(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockLoggerMockRecorder) Info(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), varargs...)
}

// Log mocks base method
func (m *MockLogger) Log(level logger.Level, format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{level, format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log
func (mr *MockLoggerMockRecorder) Log(level, format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{level, format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogger)(nil).Log), varargs...)
}

// Notice mocks base method
func (m *MockLogger) Notice(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Notice", varargs...)
}

// Notice indicates an expected call of Notice
func (mr *MockLoggerMockRecorder) Notice(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notice", reflect.TypeOf((*MockLogger)(nil).Notice), varargs...)
}

// Panic mocks base method
func (m *MockLogger) Panic(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Panic", varargs...)
}

// Panic indicates an expected call of Panic
func (mr *MockLoggerMockRecorder) Panic(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panic", reflect.TypeOf((*MockLogger)(nil).Panic), varargs...)
}

// Warning mocks base method
func (m *MockLogger) Warning(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warning", varargs...)
}

// Warning indicates an expected call of Warning
func (mr *MockLoggerMockRecorder) Warning(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockLogger)(nil).Warning), varargs...)
}

// AddHook mocks base method
func (m *MockLogger) AddHook(hook logger.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddHook", hook)
}

// AddHook indicates an expected call of AddHook
func (mr *MockLoggerMockRecorder) AddHook(hook interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHook", reflect.TypeOf((*MockLogger)(nil).AddHook), hook)
}

// With mocks base method
func (m *MockLogger) With(fields ...logger.Field) logger.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "With", varargs...)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// With indicates an expected call of With
func (mr *MockLoggerMockRecorder) With(fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockLogger)(nil).With), fields...)
}

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
func (m *MockFieldLogger) Critical(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Critical", varargs...)
}

// Critical indicates an expected call of Critical
func (mr *MockFieldLoggerMockRecorder) Critical(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Critical", reflect.TypeOf((*MockFieldLogger)(nil).Critical), varargs...)
}

// Debug mocks base method
func (m *MockFieldLogger) Debug(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockFieldLoggerMockRecorder) Debug(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockFieldLogger)(nil).Debug), varargs...)
}

// Fatal mocks base method
func (m *MockFieldLogger) Fatal(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal
func (mr *MockFieldLoggerMockRecorder) Fatal(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockFieldLogger)(nil).Fatal), varargs...)
}

// Error mocks base method
func (m *MockFieldLogger) Error(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockFieldLoggerMockRecorder) Error(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockFieldLogger)(nil).Error), varargs...)
}

// Info mocks base method
func (m *MockFieldLogger) Info(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockFieldLoggerMockRecorder) Info(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockFieldLogger)(nil).Info), varargs...)
}

// Log mocks base method
func (m *MockFieldLogger) Log(level logger.Level, format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{level, format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Log", varargs...)
}

// Log indicates an expected call of Log
func (mr *MockFieldLoggerMockRecorder) Log(level, format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{level, format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockFieldLogger)(nil).Log), varargs...)
}

// Notice mocks base method
func (m *MockFieldLogger) Notice(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Notice", varargs...)
}

// Notice indicates an expected call of Notice
func (mr *MockFieldLoggerMockRecorder) Notice(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notice", reflect.TypeOf((*MockFieldLogger)(nil).Notice), varargs...)
}

// Panic mocks base method
func (m *MockFieldLogger) Panic(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Panic", varargs...)
}

// Panic indicates an expected call of Panic
func (mr *MockFieldLoggerMockRecorder) Panic(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Panic", reflect.TypeOf((*MockFieldLogger)(nil).Panic), varargs...)
}

// Warning mocks base method
func (m *MockFieldLogger) Warning(format string, opts ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warning", varargs...)
}

// Warning indicates an expected call of Warning
func (mr *MockFieldLoggerMockRecorder) Warning(format interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockFieldLogger)(nil).Warning), varargs...)
}