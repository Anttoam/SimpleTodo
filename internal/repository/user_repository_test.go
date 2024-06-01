package repository

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/Anttoam/SimpleTodo/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	user := &domain.User{
		ID:       1,
		Name:     "test",
		Email:    "test@test.com",
		Password: "password",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Name, user.Email, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))

	ur := NewUserRepository(db)

	err = ur.Create(context.TODO(), user)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	user := &domain.User{
		ID:       1,
		Name:     "test",
		Email:    "test@test.com",
		Password: "password",
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"),
	).WithArgs(user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
			AddRow(user.ID, user.Name, user.Email, user.Password, time.Now(), time.Now()))

	ur := NewUserRepository(db)

	err = ur.FindByEmail(context.TODO(), user.Email, user)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUserRepository_FindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	user := &domain.User{
		ID:       1,
		Name:     "test",
		Email:    "test@test.com",
		Password: "password",
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?")).
		WithArgs(user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
			AddRow(user.ID, user.Name, user.Email, user.Password, time.Now(), time.Now()))

	ur := NewUserRepository(db)

	_, err = ur.FindByID(context.TODO(), user.ID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestUserRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	user := &domain.User{
		ID:       1,
		Name:     "update test",
		Email:    "update@test.com",
		Password: "updatePassword",
	}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?")).
		WithArgs(user.Name, user.Email, user.Password, sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ur := NewUserRepository(db)

	err = ur.Update(context.TODO(), user, user.ID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
