package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Anttoam/SimpleTodo/domain"
	"github.com/Anttoam/SimpleTodo/dto"
	"github.com/Anttoam/SimpleTodo/internal/usecase"
	"github.com/Anttoam/SimpleTodo/internal/usecase/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Todo"), mock.AnythingOfType("int")).Return(nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		req := dto.CreateTodoRequest{
			Title:       "test",
			Description: "test",
			UserID:      1,
		}
		err := u.Create(context.TODO(), req)
		require.NoError(t, err)
	})

	t.Run("failed_to_create_todo", func(t *testing.T) {
		mockTodoRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Todo"), mock.AnythingOfType("int")).Return(errors.New("failed to create todo")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		req := dto.CreateTodoRequest{
			Title:       "test",
			Description: "test",
			UserID:      1,
		}
		err := u.Create(context.TODO(), req)
		require.Error(t, err)
	})
}

func TestFindAll(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodos := []*domain.Todo{
		{
			ID:          1,
			Title:       "test 1",
			Description: "test 1",
			UserID:      1,
		},
		{
			ID:          2,
			Title:       "test 2",
			Description: "test 2",
			UserID:      1,
		},
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("FindAll", mock.Anything, mock.AnythingOfType("int")).Return(mockTodos, nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		userID := 1
		_, err := u.FindAll(context.TODO(), userID)
		require.NoError(t, err)
	})

	t.Run("failed_to_find_all_todo", func(t *testing.T) {
		mockTodoRepo.On("FindAll", mock.Anything, mock.AnythingOfType("int")).Return(nil, errors.New("failed to find all todo")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		userID := 1
		_, err := u.FindAll(context.TODO(), userID)
		require.Error(t, err)
	})
}

func TestFindByID(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := &domain.Todo{
		ID:          1,
		Title:       "test",
		Description: "test",
		UserID:      1,
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(mockTodo, nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		todoID := 1
		_, err := u.FindByID(context.TODO(), todoID)
		require.NoError(t, err)
	})

	t.Run("failed_to_find_todo", func(t *testing.T) {
		mockTodoRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(nil, errors.New("failed to find todo")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		todoID := 1
		_, err := u.FindByID(context.TODO(), todoID)
		require.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Todo"), mock.AnythingOfType("int")).Return(nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		req := dto.UpdateTodoRequest{
			ID:          1,
			Title:       "updated",
			Description: "updated",
		}
		err := u.Update(context.TODO(), req)
		require.NoError(t, err)
	})

	t.Run("failed_to_update_todo", func(t *testing.T) {
		mockTodoRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Todo"), mock.AnythingOfType("int")).Return(errors.New("failed to update todo")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		req := dto.UpdateTodoRequest{
			ID:          1,
			Title:       "updated",
			Description: "updated",
		}
		err := u.Update(context.TODO(), req)
		require.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		todoID := 1
		err := u.Delete(context.TODO(), todoID)
		require.NoError(t, err)
	})

	t.Run("failed_to_delete_todo", func(t *testing.T) {
		mockTodoRepo.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("failed to delete todo")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		todoID := 1
		err := u.Delete(context.TODO(), todoID)
		require.Error(t, err)
	})
}

func TestIsDone(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("UpdateDoneStatus", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("bool")).Return(nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		todoID := 1
		err := u.IsDone(context.TODO(), todoID)
		require.NoError(t, err)
	})

	t.Run("failed_to_update_done_status", func(t *testing.T) {
		mockTodoRepo.On("UpdateDoneStatus", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("bool")).Return(errors.New("failed to update done status")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		todoID := 1
		err := u.IsDone(context.TODO(), todoID)
		require.Error(t, err)
	})
}

func TestIsNotDone(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("UpdateDoneStatus", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("bool")).Return(nil).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		todoID := 1
		err := u.IsNotDone(context.TODO(), todoID)
		require.NoError(t, err)
	})

	t.Run("failed_to_update_done_status", func(t *testing.T) {
		mockTodoRepo.On("UpdateDoneStatus", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("bool")).Return(errors.New("failed to update done status")).Once()
		u := usecase.NewTodoUsecase(mockTodoRepo)
		todoID := 1
		err := u.IsNotDone(context.TODO(), todoID)
		require.Error(t, err)
	})
}
