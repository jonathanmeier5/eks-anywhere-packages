// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/authenticator/authenticator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthenticator is a mock of Authenticator interface.
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator.
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance.
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// AddSecretToAllNamespace mocks base method.
func (m *MockAuthenticator) AddSecretToAllNamespace(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecretToAllNamespace", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSecretToAllNamespace indicates an expected call of AddSecretToAllNamespace.
func (mr *MockAuthenticatorMockRecorder) AddSecretToAllNamespace(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecretToAllNamespace", reflect.TypeOf((*MockAuthenticator)(nil).AddSecretToAllNamespace), ctx)
}

// AddToConfigMap mocks base method.
func (m *MockAuthenticator) AddToConfigMap(ctx context.Context, name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToConfigMap", ctx, name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToConfigMap indicates an expected call of AddToConfigMap.
func (mr *MockAuthenticatorMockRecorder) AddToConfigMap(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToConfigMap", reflect.TypeOf((*MockAuthenticator)(nil).AddToConfigMap), ctx, name, namespace)
}

// AuthFilename mocks base method.
func (m *MockAuthenticator) AuthFilename() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthFilename")
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthFilename indicates an expected call of AuthFilename.
func (mr *MockAuthenticatorMockRecorder) AuthFilename() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthFilename", reflect.TypeOf((*MockAuthenticator)(nil).AuthFilename))
}

// DelFromConfigMap mocks base method.
func (m *MockAuthenticator) DelFromConfigMap(ctx context.Context, name, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelFromConfigMap", ctx, name, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelFromConfigMap indicates an expected call of DelFromConfigMap.
func (mr *MockAuthenticatorMockRecorder) DelFromConfigMap(ctx, name, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelFromConfigMap", reflect.TypeOf((*MockAuthenticator)(nil).DelFromConfigMap), ctx, name, namespace)
}

// GetSecretValues mocks base method.
func (m *MockAuthenticator) GetSecretValues(ctx context.Context, namespace string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValues", ctx, namespace)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretValues indicates an expected call of GetSecretValues.
func (mr *MockAuthenticatorMockRecorder) GetSecretValues(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValues", reflect.TypeOf((*MockAuthenticator)(nil).GetSecretValues), ctx, namespace)
}

// Initialize mocks base method.
func (m *MockAuthenticator) Initialize(clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockAuthenticatorMockRecorder) Initialize(clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockAuthenticator)(nil).Initialize), clusterName)
}
