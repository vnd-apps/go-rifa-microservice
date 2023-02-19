// Code generated by MockGen. DO NOT EDIT.
// Source: ports.go

// Package mock_shared is a generated GoMock package.
package mock_shared

import (
	reflect "reflect"

	shared "github.com/evmartinelli/go-rifa-microservice/internal/core/shared"
	gomock "github.com/golang/mock/gomock"
)

// MockUUIDGenerator is a mock of UUIDGenerator interface.
type MockUUIDGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockUUIDGeneratorMockRecorder
}

// MockUUIDGeneratorMockRecorder is the mock recorder for MockUUIDGenerator.
type MockUUIDGeneratorMockRecorder struct {
	mock *MockUUIDGenerator
}

// NewMockUUIDGenerator creates a new mock instance.
func NewMockUUIDGenerator(ctrl *gomock.Controller) *MockUUIDGenerator {
	mock := &MockUUIDGenerator{ctrl: ctrl}
	mock.recorder = &MockUUIDGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUUIDGenerator) EXPECT() *MockUUIDGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockUUIDGenerator) Generate() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(string)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockUUIDGeneratorMockRecorder) Generate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockUUIDGenerator)(nil).Generate))
}

// MockSlugGenerator is a mock of SlugGenerator interface.
type MockSlugGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockSlugGeneratorMockRecorder
}

// MockSlugGeneratorMockRecorder is the mock recorder for MockSlugGenerator.
type MockSlugGeneratorMockRecorder struct {
	mock *MockSlugGenerator
}

// NewMockSlugGenerator creates a new mock instance.
func NewMockSlugGenerator(ctrl *gomock.Controller) *MockSlugGenerator {
	mock := &MockSlugGenerator{ctrl: ctrl}
	mock.recorder = &MockSlugGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlugGenerator) EXPECT() *MockSlugGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockSlugGenerator) Generate(text string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", text)
	ret0, _ := ret[0].(string)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockSlugGeneratorMockRecorder) Generate(text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockSlugGenerator)(nil).Generate), text)
}

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

// CheckIsValid mocks base method.
func (m *MockAuth) CheckIsValid(bearerToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIsValid", bearerToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckIsValid indicates an expected call of CheckIsValid.
func (mr *MockAuthMockRecorder) CheckIsValid(bearerToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsValid", reflect.TypeOf((*MockAuth)(nil).CheckIsValid), bearerToken)
}

// Claims mocks base method.
func (m *MockAuth) Claims(bearerToken string) (*shared.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Claims", bearerToken)
	ret0, _ := ret[0].(*shared.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Claims indicates an expected call of Claims.
func (mr *MockAuthMockRecorder) Claims(bearerToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Claims", reflect.TypeOf((*MockAuth)(nil).Claims), bearerToken)
}

// ExtractToken mocks base method.
func (m *MockAuth) ExtractToken(bearerToken string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractToken", bearerToken)
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtractToken indicates an expected call of ExtractToken.
func (mr *MockAuthMockRecorder) ExtractToken(bearerToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractToken", reflect.TypeOf((*MockAuth)(nil).ExtractToken), bearerToken)
}
