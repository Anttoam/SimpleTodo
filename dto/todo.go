package dto

import (
	"time"

	"github.com/Anttoam/SimpleTodo/domain"
)

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=30" example:"test"`
	Description string `json:"description" validate:"max=255" example:"test"`
	UserID      int    `json:"user_id" validate:"required" example:"1"`
}

type UpdateTodoRequest struct {
	ID          int       `json:"id" validate:"required" example:"1"`
	Title       string    `json:"title" validate:"min=1,max=30" example:"updated test"`
	Description string    `json:"description" validate:"max=255" example:"updated test"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FindAllTodoResponse struct {
	Todos []*domain.Todo `json:"todos"`
}

type FindByIDTodoResponse struct {
	Todo *domain.Todo `json:"todo"`
}
