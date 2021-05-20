// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v4/pkg/domain/meta (interfaces: Author)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthor is a mock of Author interface
type MockAuthor struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorMockRecorder
}

// MockAuthorMockRecorder is the mock recorder for MockAuthor
type MockAuthorMockRecorder struct {
	mock *MockAuthor
}

// NewMockAuthor creates a new mock instance
func NewMockAuthor(ctrl *gomock.Controller) *MockAuthor {
	mock := &MockAuthor{ctrl: ctrl}
	mock.recorder = &MockAuthorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthor) EXPECT() *MockAuthorMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockAuthor) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockAuthorMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockAuthor)(nil).GetName))
}

// GetURL mocks base method
func (m *MockAuthor) GetURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetURL indicates an expected call of GetURL
func (mr *MockAuthorMockRecorder) GetURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockAuthor)(nil).GetURL))
}