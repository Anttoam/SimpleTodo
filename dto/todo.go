package dto

import (
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
)

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
}

type CreateTodoResponse struct {
	Title       string    `json:"title"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTodoRequest struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTodoResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FindAllTodoResponse struct {
	Todos []*domain.Todo `json:"todos"`
}

type FindByIDTodoResponse struct {
	Todo *domain.Todo `json:"todo"`
}
