package dto

import (
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
)

type CreateTodoRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"user_id"`
}

type CreateTodoResponse struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateTodoRequest struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateTodoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindAllTodoResponse struct {
	Todos []*domain.Todo `json:"todos"`
}

type FindByIDTodoResponse struct {
	Todo *domain.Todo `json:"todo"`
}
