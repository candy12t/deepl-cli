// Code generated by MockGen. DO NOT EDIT.
// Source: translate.go

// Package mock_webapi is a generated GoMock package.
package mock_webapi

import (
	reflect "reflect"

	entity "github.com/candy12t/deepl-cli/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockTranslater is a mock of Translater interface.
type MockTranslater struct {
	ctrl     *gomock.Controller
	recorder *MockTranslaterMockRecorder
}

// MockTranslaterMockRecorder is the mock recorder for MockTranslater.
type MockTranslaterMockRecorder struct {
	mock *MockTranslater
}

// NewMockTranslater creates a new mock instance.
func NewMockTranslater(ctrl *gomock.Controller) *MockTranslater {
	mock := &MockTranslater{ctrl: ctrl}
	mock.recorder = &MockTranslaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslater) EXPECT() *MockTranslaterMockRecorder {
	return m.recorder
}

// Translate mocks base method.
func (m *MockTranslater) Translate(arg0 *entity.Translation) (*entity.Translation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", arg0)
	ret0, _ := ret[0].(*entity.Translation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Translate indicates an expected call of Translate.
func (mr *MockTranslaterMockRecorder) Translate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockTranslater)(nil).Translate), arg0)
}
