package usecase

import (
	"context"
	"fmt"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/pkg/utils"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string, user *domain.User) error
}

type UserUsecase struct {
	ur UserRepository
}

func NewUserUsecase(ur UserRepository) *UserUsecase {
	return &UserUsecase{ur: ur}
}

func (u *UserUsecase) SignUp(ctx context.Context, req dto.SignUpRequest) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := u.ur.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user := domain.User{}
	if err := u.ur.GetUserByEmail(ctx, req.Email, &user); err != nil {
		return nil, err
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return nil, fmt.Errorf("password does not match: %w", err)
	}

	res := &dto.LoginResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	return res, nil
}
