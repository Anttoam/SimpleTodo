package usecase

import (
	"context"
	"time"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/Anttoam/golang-htmx-todos/dto"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo, userID int) error
	FindAll(ctx context.Context, userID int) ([]*domain.Todo, error)
	FindByID(ctx context.Context, id int) (*domain.Todo, error)
	Update(ctx context.Context, todo *domain.Todo, todoID int) error
}

type TodoUsecase struct {
	tr TodoRepository
}

func NewTodoUsecase(tr TodoRepository) *TodoUsecase {
	return &TodoUsecase{tr: tr}
}

func (tu *TodoUsecase) Create(ctx context.Context, req dto.CreateTodoRequest) (*dto.CreateTodoResponse, error) {
	newTodo := &domain.Todo{
		Title:     req.Title,
		Body:      req.Body,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := tu.tr.Create(ctx, newTodo, req.UserID); err != nil {
		return nil, err
	}

	res := &dto.CreateTodoResponse{
		Title:     newTodo.Title,
		Body:      newTodo.Body,
		UserID:    newTodo.UserID,
		CreatedAt: newTodo.CreatedAt,
		UpdatedAt: newTodo.UpdatedAt,
	}

	return res, nil
}

func (tu *TodoUsecase) FindAll(ctx context.Context, userID int) (*dto.FindAllTodoResponse, error) {
	todos, err := tu.tr.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := dto.FindAllTodoResponse{
		Todos: todos,
	}

	return &res, nil
}

func (tu *TodoUsecase) FindByID(ctx context.Context, todoID int) (*dto.FindByIDTodoResponse, error) {
	todo, err := tu.tr.FindByID(ctx, todoID)
	if err != nil {
		return nil, err
	}

	res := &dto.FindByIDTodoResponse{
		Todo: todo,
	}

	return res, nil
}

func (tu *TodoUsecase) Update(ctx context.Context, req dto.UpdateTodoRequest) (*dto.UpdateTodoResponse, error) {
	updateTodo := &domain.Todo{
		Title:     req.Title,
		Body:      req.Body,
		UserID:    req.UserID,
		UpdatedAt: time.Now(),
	}

	if err := tu.tr.Update(ctx, updateTodo, req.ID); err != nil {
		return nil, err
	}

	res := &dto.UpdateTodoResponse{
		ID:        req.ID,
		Title:     updateTodo.Title,
		Body:      updateTodo.Body,
		UserID:    updateTodo.UserID,
		UpdatedAt: updateTodo.UpdatedAt,
	}

	return res, nil
}
