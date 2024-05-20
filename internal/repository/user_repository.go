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

func (ur *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := ur.db.ExecContext(ctx, query, &user.Name, &user.Email, &user.Password, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string, user *domain.User) error {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"
	row := ur.db.QueryRowContext(ctx, query, email)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindByID(ctx context.Context, userID int) (*domain.User, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?"
	row := ur.db.QueryRowContext(ctx, query, userID)
	var user domain.User
	if err := row.Scan(
		&user.ID, &user.Name, &user.Email,
		&user.Password, &user.CreatedAt, &user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Update(ctx context.Context, user *domain.User, userID int) error {
	query := "UPDATE users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?"
	_, err := ur.db.ExecContext(ctx, query, &user.Name, &user.Email, &user.Password, time.Now(), userID)
	if err != nil {
		return err
	}
	return nil
}
