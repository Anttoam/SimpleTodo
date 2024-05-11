package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, &user.Name, &user.Email, &user.Password, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string, user *domain.User) error {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"
	row := r.db.QueryRowContext(ctx, query, email)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return err
	}
	return nil
}
