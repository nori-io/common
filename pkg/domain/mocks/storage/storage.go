// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nori-io/common/v4/pkg/domain/storage (interfaces: Storage)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/nori-io/common/v4/pkg/domain/storage"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Bucket mocks base method
func (m *MockStorage) Bucket(arg0 string) (storage.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bucket", arg0)
	ret0, _ := ret[0].(storage.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bucket indicates an expected call of Bucket
func (mr *MockStorageMockRecorder) Bucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bucket", reflect.TypeOf((*MockStorage)(nil).Bucket), arg0)
}

// Close mocks base method
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// CreateBucket mocks base method
func (m *MockStorage) CreateBucket(arg0 string) (storage.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0)
	ret0, _ := ret[0].(storage.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockStorageMockRecorder) CreateBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockStorage)(nil).CreateBucket), arg0)
}

// CreateBucketIfNotExists mocks base method
func (m *MockStorage) CreateBucketIfNotExists(arg0 string) (storage.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucketIfNotExists", arg0)
	ret0, _ := ret[0].(storage.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBucketIfNotExists indicates an expected call of CreateBucketIfNotExists
func (mr *MockStorageMockRecorder) CreateBucketIfNotExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucketIfNotExists", reflect.TypeOf((*MockStorage)(nil).CreateBucketIfNotExists), arg0)
}

// DeleteBucket mocks base method
func (m *MockStorage) DeleteBucket(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket
func (mr *MockStorageMockRecorder) DeleteBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockStorage)(nil).DeleteBucket), arg0)
}
