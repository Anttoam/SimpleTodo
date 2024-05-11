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

type FindAllTodoResponse struct {
	Todos []*domain.Todo `json:"todos"`
}
