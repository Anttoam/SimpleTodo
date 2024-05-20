package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/pkg/utils"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindUserByEmail(ctx context.Context, email string, user *domain.User) error
	FindByID(ctx context.Context, userID int) (*domain.User, error)
	Update(ctx context.Context, user *domain.User, userID int) error
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
	if err := u.ur.FindUserByEmail(ctx, req.Email, &user); err != nil {
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

func (u *UserUsecase) FindByID(ctx context.Context, userID int) (*dto.FindByIDUserResponse, error) {
	user, err := u.ur.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := &dto.FindByIDUserResponse{
		User: user,
	}

	return res, nil
}

func (u *UserUsecase) EditUser(ctx context.Context, req dto.UpdateUserRequest) error {
	user, err := u.ur.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	updatedUser := domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  user.Password,
		UpdatedAt: time.Now(),
	}
	if err := u.ur.Update(ctx, &updatedUser, req.ID); err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) EditPassword(ctx context.Context, req dto.UpdatePasswordRequest) error {
	user, err := u.ur.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return fmt.Errorf("password does not match: %w", err)
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	updatedUser := domain.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		UpdatedAt: time.Now(),
	}
	if err := u.ur.Update(ctx, &updatedUser, req.ID); err != nil {
		return err
	}

	return nil
}
