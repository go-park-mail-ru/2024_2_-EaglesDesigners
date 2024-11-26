// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	reflect "reflect"

	authv1 "github.com/go-park-mail-ru/2024_2_EaglesDesigner/auth_service/internal/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuth) Authenticate(ctx context.Context, in *authv1.AuthRequest) (*authv1.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, in)
	ret0, _ := ret[0].(*authv1.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthMockRecorder) Authenticate(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuth)(nil).Authenticate), ctx, in)
}

// CreateJWT mocks base method.
func (m *MockAuth) CreateJWT(ctx context.Context, in *authv1.CreateJWTRequest) (*authv1.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJWT", ctx, in)
	ret0, _ := ret[0].(*authv1.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJWT indicates an expected call of CreateJWT.
func (mr *MockAuthMockRecorder) CreateJWT(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJWT", reflect.TypeOf((*MockAuth)(nil).CreateJWT), ctx, in)
}

// GetUserByJWT mocks base method.
func (m *MockAuth) GetUserByJWT(ctx context.Context, in *authv1.Token) (*authv1.UserJWT, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByJWT", ctx, in)
	ret0, _ := ret[0].(*authv1.UserJWT)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByJWT indicates an expected call of GetUserByJWT.
func (mr *MockAuthMockRecorder) GetUserByJWT(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByJWT", reflect.TypeOf((*MockAuth)(nil).GetUserByJWT), ctx, in)
}

// GetUserDataByUsername mocks base method.
func (m *MockAuth) GetUserDataByUsername(ctx context.Context, in *authv1.GetUserDataByUsernameRequest) (*authv1.GetUserDataByUsernameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDataByUsername", ctx, in)
	ret0, _ := ret[0].(*authv1.GetUserDataByUsernameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDataByUsername indicates an expected call of GetUserDataByUsername.
func (mr *MockAuthMockRecorder) GetUserDataByUsername(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDataByUsername", reflect.TypeOf((*MockAuth)(nil).GetUserDataByUsername), ctx, in)
}

// IsAuthorized mocks base method.
func (m *MockAuth) IsAuthorized(ctx context.Context, in *authv1.Token) (*authv1.UserJWT, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthorized", ctx, in)
	ret0, _ := ret[0].(*authv1.UserJWT)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuthorized indicates an expected call of IsAuthorized.
func (mr *MockAuthMockRecorder) IsAuthorized(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorized", reflect.TypeOf((*MockAuth)(nil).IsAuthorized), ctx, in)
}

// Registration mocks base method.
func (m *MockAuth) Registration(ctx context.Context, in *authv1.RegistrationRequest) (*authv1.Nothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Registration", ctx, in)
	ret0, _ := ret[0].(*authv1.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Registration indicates an expected call of Registration.
func (mr *MockAuthMockRecorder) Registration(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Registration", reflect.TypeOf((*MockAuth)(nil).Registration), ctx, in)
}
