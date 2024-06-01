// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/Anttoam/SimpleTodo/dto"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// EditPassword provides a mock function with given fields: ctx, user
func (_m *UserUsecase) EditPassword(ctx context.Context, user dto.UpdatePasswordRequest) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for EditPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdatePasswordRequest) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditUser provides a mock function with given fields: ctx, user
func (_m *UserUsecase) EditUser(ctx context.Context, user dto.UpdateUserRequest) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for EditUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdateUserRequest) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUserByID provides a mock function with given fields: ctx, userID
func (_m *UserUsecase) FindUserByID(ctx context.Context, userID int) (*dto.FindByIDUserResponse, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByID")
	}

	var r0 *dto.FindByIDUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*dto.FindByIDUserResponse, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *dto.FindByIDUserResponse); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.FindByIDUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, req
func (_m *UserUsecase) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *dto.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) (*dto.LoginResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) *dto.LoginResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.LoginRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: ctx, req
func (_m *UserUsecase) SignUp(ctx context.Context, req dto.SignUpRequest) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for SignUp")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.SignUpRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
