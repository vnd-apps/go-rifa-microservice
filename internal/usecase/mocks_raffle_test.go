// Code generated by MockGen. DO NOT EDIT.
// Source: raffle_interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"

	entity "github.com/evmartinelli/go-rifa-microservice/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockRaffle is a mock of Raffle interface.
type MockRaffle struct {
	ctrl     *gomock.Controller
	recorder *MockRaffleMockRecorder
}

// MockRaffleMockRecorder is the mock recorder for MockRaffle.
type MockRaffleMockRecorder struct {
	mock *MockRaffle
}

// NewMockRaffle creates a new mock instance.
func NewMockRaffle(ctrl *gomock.Controller) *MockRaffle {
	mock := &MockRaffle{ctrl: ctrl}
	mock.recorder = &MockRaffleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRaffle) EXPECT() *MockRaffleMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRaffle) Create(arg0 context.Context, arg1 entity.Raffle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRaffleMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRaffle)(nil).Create), arg0, arg1)
}

// GetAvailableRaffle mocks base method.
func (m *MockRaffle) GetAvailableRaffle(arg0 context.Context) ([]entity.Raffle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableRaffle", arg0)
	ret0, _ := ret[0].([]entity.Raffle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableRaffle indicates an expected call of GetAvailableRaffle.
func (mr *MockRaffleMockRecorder) GetAvailableRaffle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableRaffle", reflect.TypeOf((*MockRaffle)(nil).GetAvailableRaffle), arg0)
}

// MockRaffleRepo is a mock of RaffleRepo interface.
type MockRaffleRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRaffleRepoMockRecorder
}

// MockRaffleRepoMockRecorder is the mock recorder for MockRaffleRepo.
type MockRaffleRepoMockRecorder struct {
	mock *MockRaffleRepo
}

// NewMockRaffleRepo creates a new mock instance.
func NewMockRaffleRepo(ctrl *gomock.Controller) *MockRaffleRepo {
	mock := &MockRaffleRepo{ctrl: ctrl}
	mock.recorder = &MockRaffleRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRaffleRepo) EXPECT() *MockRaffleRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRaffleRepo) Create(arg0 context.Context, arg1 entity.Raffle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRaffleRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRaffleRepo)(nil).Create), arg0, arg1)
}

// GetAvailableRaffle mocks base method.
func (m *MockRaffleRepo) GetAvailableRaffle(arg0 context.Context) ([]entity.Raffle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableRaffle", arg0)
	ret0, _ := ret[0].([]entity.Raffle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableRaffle indicates an expected call of GetAvailableRaffle.
func (mr *MockRaffleRepoMockRecorder) GetAvailableRaffle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableRaffle", reflect.TypeOf((*MockRaffleRepo)(nil).GetAvailableRaffle), arg0)
}
