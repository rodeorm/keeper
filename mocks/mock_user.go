// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rodeorm/keeper/internal/core (interfaces: UserStorager)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/rodeorm/keeper/internal/core"
)

// MockUserStorager is a mock of UserStorager interface.
type MockUserStorager struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoragerMockRecorder
}

// MockUserStoragerMockRecorder is the mock recorder for MockUserStorager.
type MockUserStoragerMockRecorder struct {
	mock *MockUserStorager
}

// NewMockUserStorager creates a new mock instance.
func NewMockUserStorager(ctrl *gomock.Controller) *MockUserStorager {
	mock := &MockUserStorager{ctrl: ctrl}
	mock.recorder = &MockUserStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStorager) EXPECT() *MockUserStoragerMockRecorder {
	return m.recorder
}

// AuthUser mocks base method.
func (m *MockUserStorager) AuthUser(arg0 context.Context, arg1 *core.User) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUser", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthUser indicates an expected call of AuthUser.
func (mr *MockUserStoragerMockRecorder) AuthUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUser", reflect.TypeOf((*MockUserStorager)(nil).AuthUser), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockUserStorager) DeleteUser(arg0 context.Context, arg1 *core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserStoragerMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserStorager)(nil).DeleteUser), arg0, arg1)
}

// RegUser mocks base method.
func (m *MockUserStorager) RegUser(arg0 context.Context, arg1 *core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegUser indicates an expected call of RegUser.
func (mr *MockUserStoragerMockRecorder) RegUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegUser", reflect.TypeOf((*MockUserStorager)(nil).RegUser), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockUserStorager) UpdateUser(arg0 context.Context, arg1 *core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserStoragerMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserStorager)(nil).UpdateUser), arg0, arg1)
}

// VerifyUserOTP mocks base method.
func (m *MockUserStorager) VerifyUserOTP(arg0 context.Context, arg1 int, arg2 *core.User) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUserOTP", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyUserOTP indicates an expected call of VerifyUserOTP.
func (mr *MockUserStoragerMockRecorder) VerifyUserOTP(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUserOTP", reflect.TypeOf((*MockUserStorager)(nil).VerifyUserOTP), arg0, arg1, arg2)
}
