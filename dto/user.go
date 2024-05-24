package dto

import (
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
)

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignUpResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type FindByIDUserResponse struct {
	User *domain.User `json:"user"`
}

type UpdateUserRequest struct {
	ID        int       `json:"id" validate:"required"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePasswordRequest struct {
	ID          int       `json:"id" validate:"required"`
	Password    string    `json:"password" validate:"required,min=6"`
	NewPassword string    `json:"new_password" validate:"required,min=6"`
	UpdatedAt   time.Time `json:"updated_at"`
}
