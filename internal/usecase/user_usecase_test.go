package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/internal/usecase"
	"github.com/Anttoam/golang-htmx-todos/internal/usecase/mocks"
	"github.com/Anttoam/golang-htmx-todos/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
		u := usecase.NewUserUsecase(mockUserRepo)
		req := dto.SignUpRequest{
			Name:     "test",
			Email:    "test@example.com",
			Password: "password",
		}
		err := u.SignUp(context.TODO(), req)
		require.NoError(t, err)
	})

	t.Run("failed_to_create_user", func(t *testing.T) {
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(errors.New("failed to create user")).Once()
		u := usecase.NewUserUsecase(mockUserRepo)
		req := dto.SignUpRequest{
			Name:     "test",
			Email:    "test@example.com",
			Password: "password",
		}
		err := u.SignUp(context.TODO(), req)
		require.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	hashedPassword, err := utils.HashPassword("password")
	require.NoError(t, err)

	user := domain.User{
		ID:       1,
		Name:     "test",
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("*domain.User")).Run(func(args mock.Arguments) {
			argUser := args.Get(2).(*domain.User)
			*argUser = user
		}).Return(nil).Once()

		u := usecase.NewUserUsecase(mockUserRepo)
		req := dto.LoginRequest{
			Email:    "test@example.com",
			Password: "password",
		}
		_, err = u.Login(context.TODO(), req)
		require.NoError(t, err)
	})
}

func TestFindUserByID(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &domain.User{
		ID:       1,
		Name:     "test",
		Email:    "test@example.com",
		Password: "password",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(mockUser, nil).Once()
		u := usecase.NewUserUsecase(mockUserRepo)
		userID := 1
		_, err := u.FindUserByID(context.TODO(), userID)
		require.NoError(t, err)
	})

	t.Run("failed_to_find_user", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(nil, errors.New("user not found")).Once()
		u := usecase.NewUserUsecase(mockUserRepo)
		userID := 1
		_, err := u.FindUserByID(context.TODO(), userID)
		require.Error(t, err)
	})
}

func TestEditUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &domain.User{
		ID:       1,
		Name:     "test",
		Email:    "test@example.com",
		Password: "password",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(mockUser, nil).Once()
		mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User"), mock.AnythingOfType("int")).Return(nil).Once()
		u := usecase.NewUserUsecase(mockUserRepo)
		req := dto.UpdateUserRequest{
			Name:  "updated",
			Email: "updated@example.com",
		}
		err := u.EditUser(context.TODO(), req)
		require.NoError(t, err)
	})

	t.Run("Update error", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(mockUser, nil).Once()
		mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User"), mock.AnythingOfType("int")).Return(errors.New("update failed")).Once()

		u := usecase.NewUserUsecase(mockUserRepo)
		req := dto.UpdateUserRequest{
			ID:    1,
			Name:  "new name",
			Email: "newemail@example.com",
		}

		err := u.EditUser(context.TODO(), req)
		require.Error(t, err)
		require.Equal(t, "update failed", err.Error())
	})

}

func TestEditPassword(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	hashedPassword, err := utils.HashPassword("password")
	require.NoError(t, err)
	mockUser := &domain.User{
		ID:       1,
		Name:     "test",
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(mockUser, nil).Once()
		mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User"), mock.AnythingOfType("int")).Return(nil).Once()
		u := usecase.NewUserUsecase(mockUserRepo)
		require.NoError(t, err)
		req := dto.UpdatePasswordRequest{
			Password:    "password",
			NewPassword: "newpassword",
		}
		err = u.EditPassword(context.TODO(), req)
		require.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mockUserRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(mockUser, nil).Once()
		mockUserRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User"), mock.AnythingOfType("int")).Return(errors.New("update failed")).Once()

		u := usecase.NewUserUsecase(mockUserRepo)
		req := dto.UpdatePasswordRequest{
			Password:    "password",
			NewPassword: "newpassword",
		}
		err := u.EditPassword(context.TODO(), req)
		require.Error(t, err)
		require.Equal(t, "update failed", err.Error())
	})
}
