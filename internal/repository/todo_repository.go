package repository

import (
	"context"
	"database/sql"
	"log"

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

func (r *TodoRepository) FindAll(ctx context.Context, userID int) ([]*domain.Todo, error) {
	query := "SELECT id, title, body, user_id, created_at, updated_at FROM todos WHERE user_id = ?"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var todos []*domain.Todo
	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Body, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}
