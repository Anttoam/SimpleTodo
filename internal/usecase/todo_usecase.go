package usecase

import (
	"context"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/Anttoam/golang-htmx-todos/dto"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo, userID int) error
}

type TodoUsecase struct {
	tr TodoRepository
}

func NewTodoUsecase(tr TodoRepository) *TodoUsecase {
	return &TodoUsecase{tr: tr}
}

func (tu *TodoUsecase) Create(ctx context.Context, req dto.CreateTodoRequest) (*dto.CreateTodoResponse, error) {
	newTodo := domain.Todo{
		Title:  req.Title,
		Body:   req.Body,
		UserID: req.UserID,
	}
	if err := tu.tr.Create(ctx, &newTodo, req.UserID); err != nil {
		return nil, err
	}

	res := dto.CreateTodoResponse{
		Title:     newTodo.Title,
		Body:      newTodo.Body,
		UserID:    newTodo.UserID,
		CreatedAt: newTodo.CreatedAt,
	}

	return &res, nil
}
