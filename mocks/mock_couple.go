// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rodeorm/keeper/internal/core (interfaces: CoupleStorager)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/rodeorm/keeper/internal/core"
)

// MockCoupleStorager is a mock of CoupleStorager interface.
type MockCoupleStorager struct {
	ctrl     *gomock.Controller
	recorder *MockCoupleStoragerMockRecorder
}

// MockCoupleStoragerMockRecorder is the mock recorder for MockCoupleStorager.
type MockCoupleStoragerMockRecorder struct {
	mock *MockCoupleStorager
}

// NewMockCoupleStorager creates a new mock instance.
func NewMockCoupleStorager(ctrl *gomock.Controller) *MockCoupleStorager {
	mock := &MockCoupleStorager{ctrl: ctrl}
	mock.recorder = &MockCoupleStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoupleStorager) EXPECT() *MockCoupleStoragerMockRecorder {
	return m.recorder
}

// AddCoupleByUser mocks base method.
func (m *MockCoupleStorager) AddCoupleByUser(arg0 context.Context, arg1 *core.Couple, arg2 *core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCoupleByUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCoupleByUser indicates an expected call of AddCoupleByUser.
func (mr *MockCoupleStoragerMockRecorder) AddCoupleByUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCoupleByUser", reflect.TypeOf((*MockCoupleStorager)(nil).AddCoupleByUser), arg0, arg1, arg2)
}

// DeleteCoupleByUser mocks base method.
func (m *MockCoupleStorager) DeleteCoupleByUser(arg0 context.Context, arg1 *core.Couple, arg2 *core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCoupleByUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCoupleByUser indicates an expected call of DeleteCoupleByUser.
func (mr *MockCoupleStoragerMockRecorder) DeleteCoupleByUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCoupleByUser", reflect.TypeOf((*MockCoupleStorager)(nil).DeleteCoupleByUser), arg0, arg1, arg2)
}

// SelectAllCouplesByUser mocks base method.
func (m *MockCoupleStorager) SelectAllCouplesByUser(arg0 context.Context, arg1 *core.User) ([]core.Couple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAllCouplesByUser", arg0, arg1)
	ret0, _ := ret[0].([]core.Couple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAllCouplesByUser indicates an expected call of SelectAllCouplesByUser.
func (mr *MockCoupleStoragerMockRecorder) SelectAllCouplesByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAllCouplesByUser", reflect.TypeOf((*MockCoupleStorager)(nil).SelectAllCouplesByUser), arg0, arg1)
}

// UpdateCoupleByUser mocks base method.
func (m *MockCoupleStorager) UpdateCoupleByUser(arg0 context.Context, arg1 *core.Couple, arg2 *core.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCoupleByUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCoupleByUser indicates an expected call of UpdateCoupleByUser.
func (mr *MockCoupleStoragerMockRecorder) UpdateCoupleByUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCoupleByUser", reflect.TypeOf((*MockCoupleStorager)(nil).UpdateCoupleByUser), arg0, arg1, arg2)
}
