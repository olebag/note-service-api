// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scipie28/test/internal/saver (interfaces: Saver)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/scipie28/note-service-api/internal/app/model"
)

// MockSaver is a mock of Saver interface.
type MockSaver struct {
	ctrl     *gomock.Controller
	recorder *MockSaverMockRecorder
}

// MockSaverMockRecorder is the mock recorder for MockSaver.
type MockSaverMockRecorder struct {
	mock *MockSaver
}

// NewMockSaver creates a new mock instance.
func NewMockSaver(ctrl *gomock.Controller) *MockSaver {
	mock := &MockSaver{ctrl: ctrl}
	mock.recorder = &MockSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaver) EXPECT() *MockSaverMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSaver) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSaverMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSaver)(nil).Close))
}

// Init mocks base method.
func (m *MockSaver) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockSaverMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSaver)(nil).Init))
}

// Save mocks base method.
func (m *MockSaver) Save(arg0 api.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockSaverMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockSaver)(nil).Save), arg0)
}
