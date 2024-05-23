package usecase

import (
	"context"
	"testing"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/Anttoam/golang-htmx-todos/dto"
	mock_usecase "github.com/Anttoam/golang-htmx-todos/internal/usecase/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestTodoCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock_usecase.NewMockTodoRepository(ctrl)
	tu := NewTodoUsecase(tr)

	ctx := context.Background()
	tr.EXPECT().Create(ctx, gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	req := dto.CreateTodoRequest{
		Title:       "test",
		Description: "test",
		UserID:      1,
	}

	err := tu.Create(ctx, req)
	require.NoError(t, err)
}

func TestTodoFindAll(t *testing.T) {
	user, _ := RandomUser(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock_usecase.NewMockTodoRepository(ctrl)
	tu := NewTodoUsecase(tr)

	ctx := context.Background()
	tr.EXPECT().FindAll(ctx, gomock.Any()).
		Times(1).
		DoAndReturn(func(ctx context.Context, userID int) ([]*domain.Todo, error) {
			return []*domain.Todo{
				{
					ID:          user.ID,
					Title:       "test1",
					Description: "test",
					UserID:      1,
				},
				{
					ID:          user.ID,
					Title:       "test2",
					Description: "test",
					UserID:      1,
				},
			}, nil
		})

	_, err := tu.FindAll(ctx, user.ID)
	require.NoError(t, err)
}

func TestTodoFindByID(t *testing.T) {
	user, _ := RandomUser(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock_usecase.NewMockTodoRepository(ctrl)
	tu := NewTodoUsecase(tr)

	ctx := context.Background()
	tr.EXPECT().FindByID(ctx, gomock.Any()).
		Times(1).
		DoAndReturn(func(ctx context.Context, id int) (*domain.Todo, error) {
			return &domain.Todo{
				ID:          1,
				Title:       "test",
				Description: "test",
				UserID:      user.ID,
			}, nil
		})

	_, err := tu.FindByID(ctx, 1)
	require.NoError(t, err)
}

func TestTodoUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock_usecase.NewMockTodoRepository(ctrl)
	tu := NewTodoUsecase(tr)

	ctx := context.Background()
	tr.EXPECT().Update(ctx, gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	req := dto.UpdateTodoRequest{
		ID:          1,
		Title:       "test",
		Description: "test",
	}

	err := tu.Update(ctx, req)
	require.NoError(t, err)
}

func TestTodoDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock_usecase.NewMockTodoRepository(ctrl)
	tu := NewTodoUsecase(tr)

	ctx := context.Background()
	tr.EXPECT().Delete(ctx, gomock.Any()).
		Times(1).
		Return(nil)

	err := tu.Delete(ctx, 1)
	require.NoError(t, err)
}

func TestTodoIsDone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock_usecase.NewMockTodoRepository(ctrl)
	tu := NewTodoUsecase(tr)

	ctx := context.Background()
	tr.EXPECT().UpdateDoneStatus(ctx, gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	err := tu.IsDone(ctx, 1)
	require.NoError(t, err)
}

func TestTodoIsNotDone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock_usecase.NewMockTodoRepository(ctrl)
	tu := NewTodoUsecase(tr)

	ctx := context.Background()
	tr.EXPECT().UpdateDoneStatus(ctx, gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	err := tu.IsNotDone(ctx, 1)
	require.NoError(t, err)
}
