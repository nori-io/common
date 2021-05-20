// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v5/pkg/domain/version (interfaces: Constraints)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	version "github.com/nori-io/common/v5/pkg/domain/version"
)

// MockConstraints is a mock of Constraints interface.
type MockConstraints struct {
	ctrl     *gomock.Controller
	recorder *MockConstraintsMockRecorder
}

// MockConstraintsMockRecorder is the mock recorder for MockConstraints.
type MockConstraintsMockRecorder struct {
	mock *MockConstraints
}

// NewMockConstraints creates a new mock instance.
func NewMockConstraints(ctrl *gomock.Controller) *MockConstraints {
	mock := &MockConstraints{ctrl: ctrl}
	mock.recorder = &MockConstraintsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConstraints) EXPECT() *MockConstraintsMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockConstraints) Check(arg0 version.Version) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockConstraintsMockRecorder) Check(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockConstraints)(nil).Check), arg0)
}

// String mocks base method.
func (m *MockConstraints) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockConstraintsMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockConstraints)(nil).String))
}
