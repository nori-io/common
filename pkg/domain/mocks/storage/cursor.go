// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v4/pkg/domain/storage (interfaces: Cursor)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCursor is a mock of Cursor interface
type MockCursor struct {
	ctrl     *gomock.Controller
	recorder *MockCursorMockRecorder
}

// MockCursorMockRecorder is the mock recorder for MockCursor
type MockCursorMockRecorder struct {
	mock *MockCursor
}

// NewMockCursor creates a new mock instance
func NewMockCursor(ctrl *gomock.Controller) *MockCursor {
	mock := &MockCursor{ctrl: ctrl}
	mock.recorder = &MockCursorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCursor) EXPECT() *MockCursorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCursor) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockCursorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCursor)(nil).Close))
}

// Delete mocks base method
func (m *MockCursor) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCursorMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCursor)(nil).Delete))
}

// First mocks base method
func (m *MockCursor) First() (string, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "First")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// First indicates an expected call of First
func (mr *MockCursorMockRecorder) First() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockCursor)(nil).First))
}

// HasNext mocks base method
func (m *MockCursor) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext
func (mr *MockCursorMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockCursor)(nil).HasNext))
}

// Last mocks base method
func (m *MockCursor) Last() (string, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Last indicates an expected call of Last
func (mr *MockCursorMockRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockCursor)(nil).Last))
}

// Next mocks base method
func (m *MockCursor) Next() (string, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockCursorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockCursor)(nil).Next))
}

// Prev mocks base method
func (m *MockCursor) Prev() (string, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prev")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Prev indicates an expected call of Prev
func (mr *MockCursorMockRecorder) Prev() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prev", reflect.TypeOf((*MockCursor)(nil).Prev))
}

// Seek mocks base method
func (m *MockCursor) Seek(arg0 string) (string, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seek", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Seek indicates an expected call of Seek
func (mr *MockCursorMockRecorder) Seek(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seek", reflect.TypeOf((*MockCursor)(nil).Seek), arg0)
}
