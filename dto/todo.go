package dto

import "time"

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
}
