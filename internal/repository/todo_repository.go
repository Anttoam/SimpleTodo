package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Anttoam/SimpleTodo/domain"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(ctx context.Context, todo *domain.Todo, userID int) error {
	query := "INSERT INTO todos (title, description, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, todo.Title, todo.Description, userID, todo.CreatedAt, todo.UpdatedAt)
	return err
}

func (r *TodoRepository) FindAll(ctx context.Context, userID int) ([]*domain.Todo, error) {
	query := "SELECT id, title, description, done, user_id, created_at, updated_at FROM todos WHERE user_id = ?"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	var todos []*domain.Todo
	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(
			&todo.ID, &todo.Title, &todo.Description,
			&todo.Done, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt,
		); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (r *TodoRepository) FindByID(ctx context.Context, todoID int) (*domain.Todo, error) {
	query := "SELECT id, title, description, done, user_id, created_at, updated_at FROM todos WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, todoID)
	var todo domain.Todo
	if err := row.Scan(
		&todo.ID, &todo.Title, &todo.Description,
		&todo.Done, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Update(ctx context.Context, todo *domain.Todo, todoID int) error {
	query := "UPDATE todos SET title = ?, description = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, todo.Title, todo.Description, todo.UpdatedAt, todoID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) Delete(ctx context.Context, todoID int) error {
	query := "DELETE FROM todos WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, todoID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) UpdateDoneStatus(ctx context.Context, todoID int, done bool) error {
	query := "UPDATE todos SET done = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, done, todoID)
	if err != nil {
		return err
	}
	return nil
}
