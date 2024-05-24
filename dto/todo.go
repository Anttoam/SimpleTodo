package dto

import (
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
)

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=30"`
	Description string `json:"description"`
	UserID      int    `json:"user_id" validate:"required"`
}

type UpdateTodoRequest struct {
	ID          int       `json:"id" validate:"required"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FindAllTodoResponse struct {
	Todos []*domain.Todo `json:"todos"`
}

type FindByIDTodoResponse struct {
	Todo *domain.Todo `json:"todo"`
}
