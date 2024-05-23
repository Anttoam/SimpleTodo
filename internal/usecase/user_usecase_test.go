package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/Anttoam/golang-htmx-todos/dto"
	mock_usecase "github.com/Anttoam/golang-htmx-todos/internal/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := mock_usecase.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	req := dto.SignUpRequest{
		Name:     "test",
		Email:    "test@example.com",
		Password: "password",
	}

	ctx := context.Background()
	ur.EXPECT().Create(ctx, gomock.Any()).
		Times(1).
		Return(nil)

	err := uu.SignUp(ctx, req)
	require.NoError(t, err)
}

func TestLogin(t *testing.T) {
	user, password := RandomUser(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := mock_usecase.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	req := &dto.LoginRequest{
		Email:    user.Email,
		Password: password,
	}

	ctx := context.Background()
	ur.EXPECT().FindByEmail(ctx, req.Email, gomock.Any()).Times(1).DoAndReturn(
		func(_ context.Context, _ string, u *domain.User) error {
			*u = user
			return nil
		},
	)

	res, err := uu.Login(ctx, *req)

	require.NoError(t, err)
	require.NotNil(t, res)
	assert.Equal(t, user.ID, res.ID)
	assert.Equal(t, user.Email, res.Email)
}

func TestFindUserByID(t *testing.T) {
	user, _ := RandomUser(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := mock_usecase.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	ctx := context.Background()
	ur.EXPECT().FindByID(ctx, gomock.Any()).
		Times(1).
		DoAndReturn(
			func(_ context.Context, _ int) (*domain.User, error) {
				return &user, nil
			},
		)

	res, err := uu.FindUserByID(ctx, user.ID)
	require.NoError(t, err)
	assert.Equal(t, user.ID, res.User.ID)
	assert.Equal(t, user.Name, res.User.Name)
	assert.Equal(t, user.Email, res.User.Email)
	assert.Equal(t, user.CreatedAt, res.User.CreatedAt)
	assert.Equal(t, user.UpdatedAt, res.User.UpdatedAt)
}

func TestEditUser(t *testing.T) {
	user, _ := RandomUser(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := mock_usecase.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	ctx := context.Background()

	ur.EXPECT().FindByID(ctx, gomock.Any()).
		Times(1).
		DoAndReturn(
			func(_ context.Context, _ int) (*domain.User, error) {
				return &user, nil
			},
		)

	ur.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	req := dto.UpdateUserRequest{
		Name:      "updateUser",
		Email:     "updateUser@example.com",
		UpdatedAt: time.Now(),
	}

	err := uu.EditUser(ctx, req)
	require.NoError(t, err)
}

func TestEditPassword(t *testing.T) {
	user, _ := RandomUser(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := mock_usecase.NewMockUserRepository(ctrl)
	uu := NewUserUsecase(ur)

	ctx := context.Background()

	ur.EXPECT().FindByID(ctx, gomock.Any()).
		Times(1).
		DoAndReturn(
			func(_ context.Context, _ int) (*domain.User, error) {
				return &user, nil
			},
		)

	ur.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	req := dto.UpdatePasswordRequest{
		Password:    "password",
		NewPassword: "newPassword",
		UpdatedAt:   time.Now(),
	}

	err := uu.EditPassword(ctx, req)
	require.NoError(t, err)
}
