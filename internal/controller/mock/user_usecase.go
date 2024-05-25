// Code generated by MockGen. DO NOT EDIT.
// Source: internal/controller/user_controller.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	dto "github.com/Anttoam/golang-htmx-todos/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// EditPassword mocks base method.
func (m *MockUserUsecase) EditPassword(ctx context.Context, user dto.UpdatePasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditPassword", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditPassword indicates an expected call of EditPassword.
func (mr *MockUserUsecaseMockRecorder) EditPassword(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditPassword", reflect.TypeOf((*MockUserUsecase)(nil).EditPassword), ctx, user)
}

// EditUser mocks base method.
func (m *MockUserUsecase) EditUser(ctx context.Context, user dto.UpdateUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditUser indicates an expected call of EditUser.
func (mr *MockUserUsecaseMockRecorder) EditUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUser", reflect.TypeOf((*MockUserUsecase)(nil).EditUser), ctx, user)
}

// FindUserByID mocks base method.
func (m *MockUserUsecase) FindUserByID(ctx context.Context, userID int) (*dto.FindByIDUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", ctx, userID)
	ret0, _ := ret[0].(*dto.FindByIDUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockUserUsecaseMockRecorder) FindUserByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockUserUsecase)(nil).FindUserByID), ctx, userID)
}

// Login mocks base method.
func (m *MockUserUsecase) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, req)
	ret0, _ := ret[0].(*dto.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserUsecaseMockRecorder) Login(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserUsecase)(nil).Login), ctx, req)
}

// SignUp mocks base method.
func (m *MockUserUsecase) SignUp(ctx context.Context, req dto.SignUpRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserUsecaseMockRecorder) SignUp(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUserUsecase)(nil).SignUp), ctx, req)
}
