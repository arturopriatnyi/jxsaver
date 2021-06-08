// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CreateFile mocks base method.
func (m *MockManager) CreateFile(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFile", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFile indicates an expected call of CreateFile.
func (mr *MockManagerMockRecorder) CreateFile(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFile", reflect.TypeOf((*MockManager)(nil).CreateFile), name)
}

// WriteToFile mocks base method.
func (m *MockManager) WriteToFile(fileName, content string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToFile", fileName, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToFile indicates an expected call of WriteToFile.
func (mr *MockManagerMockRecorder) WriteToFile(fileName, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToFile", reflect.TypeOf((*MockManager)(nil).WriteToFile), fileName, content)
}
