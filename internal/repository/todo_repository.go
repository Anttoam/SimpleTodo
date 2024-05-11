package repository

import (
	"context"
	"database/sql"

	"github.com/Anttoam/golang-htmx-todos/domain"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(ctx context.Context, todo *domain.Todo, userID int) error {
	query := "INSERT INTO todos (title, body, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, todo.Title, todo.Body, userID, todo.CreatedAt, todo.UpdatedAt)
	return err
}
