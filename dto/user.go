package dto

import (
	"time"

	"github.com/Anttoam/SimpleTodo/domain"
)

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,min=2" example:"testuser"`
	Email    string `json:"email" validate:"required,email" example:"testuser@test.com"`
	Password string `json:"password" validate:"required,min=6" example:"password"`
}

type SignUpResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"testuser"`
	Password string `json:"password" validate:"required,min=6" example:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type FindByIDUserResponse struct {
	User *domain.User `json:"user"`
}

type UpdateUserRequest struct {
	ID        int       `json:"id" validate:"required" example:"1"`
	Name      string    `json:"name" validate:"min=2" example:"updateduser"`
	Email     string    `json:"email" validate:"email" example:"updated@test.com"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePasswordRequest struct {
	ID          int       `json:"id" validate:"required"`
	Password    string    `json:"password" validate:"required,min=6"`
	NewPassword string    `json:"new_password" validate:"required,min=6"`
	UpdatedAt   time.Time `json:"updated_at"`
}
