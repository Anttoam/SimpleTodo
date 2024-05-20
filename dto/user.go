package dto

import (
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type FindByIDUserResponse struct {
	User *domain.User `json:"user"`
}

type UpdateUserRequest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePasswordRequest struct {
	ID          int       `json:"id"`
	Password    string    `json:"password"`
	NewPassword string    `json:"new_password"`
	UpdatedAt   time.Time `json:"updated_at"`
}
