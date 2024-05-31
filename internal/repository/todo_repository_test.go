package repository

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestTodoCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	todo := &domain.Todo{
		ID:          1,
		Title:       "test",
		Description: "test",
		UserID:      1,
	}

	mock.ExpectExec("INSERT INTO todos").
		WithArgs(todo.Title, todo.Description, todo.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(int64(todo.ID), 1))

	tr := NewTodoRepository(db)

	err = tr.Create(context.TODO(), todo, todo.UserID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestTodoFindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	userID := 1

	todo := domain.Todo{
		ID:          1,
		Title:       "test1",
		Description: "test1",
		UserID:      userID,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT id, title, description, done, user_id, created_at, updated_at FROM todos WHERE user_id = ?"),
	).WithArgs(userID).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"title",
		"description",
		"done",
		"user_id",
		"created_at",
		"updated_at",
	}).AddRow(todo.ID, todo.Title, todo.Description, todo.Done, todo.UserID, todo.CreatedAt, todo.UpdatedAt))

	tr := NewTodoRepository(db)

	_, err = tr.FindAll(context.TODO(), userID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestTodoFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	todo := &domain.Todo{
		ID:          1,
		Title:       "test",
		Description: "test",
		UserID:      1,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT id, title, description, done, user_id, created_at, updated_at FROM todos WHERE id = ?"),
	).WithArgs(todo.ID).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"title",
		"description",
		"done",
		"user_id",
		"created_at",
		"updated_at",
	}).AddRow(todo.ID, todo.Title, todo.Description, todo.Done, todo.UserID, todo.CreatedAt, todo.UpdatedAt))

	tr := NewTodoRepository(db)

	_, err = tr.FindByID(context.TODO(), todo.ID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestTodoUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	todo := &domain.Todo{
		ID:          1,
		Title:       "update",
		Description: "update",
		Done:        false,
		UserID:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE todos SET title = ?, description = ?, updated_at = ? WHERE id = ?")).
		WithArgs(todo.Title, todo.Description, sqlmock.AnyArg(), todo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	tr := NewTodoRepository(db)

	err = tr.Update(context.TODO(), todo, todo.ID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestTodoDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	todo := &domain.Todo{
		ID:          1,
		Title:       "test",
		Description: "test",
		UserID:      1,
	}

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM todos WHERE id = ?")).
		WithArgs(todo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	tr := NewTodoRepository(db)

	err = tr.Delete(context.TODO(), todo.ID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestTodoUpdateDoneStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	todo := &domain.Todo{
		ID:          1,
		Title:       "update",
		Description: "update",
		Done:        true,
		UserID:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE todos SET done = ? WHERE id = ?")).
		WithArgs(todo.Done, todo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewTodoRepository(db)

	err = repo.UpdateDoneStatus(context.TODO(), todo.ID, todo.Done)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
