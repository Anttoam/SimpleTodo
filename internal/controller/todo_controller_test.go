package controller_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/internal/controller"
	"github.com/Anttoam/golang-htmx-todos/internal/controller/mocks"
	"github.com/Anttoam/golang-htmx-todos/pkg/validation"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	mockUcase := new(mocks.TodoUsecase)
	request := dto.CreateTodoRequest{
		Title:       "test",
		Description: "test",
		UserID:      1,
	}
	j, err := json.Marshal(request)
	require.NoError(t, err)

	mockUcase.On("Create", mock.Anything, request).Return(nil).Once()
	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "todo/create", strings.NewReader(string(j)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.TodoController{
		Usecase: mockUcase,
		Store:   store,
	}

	session, _ := store.Get(req, "session_id")
	session.Values["id"] = 1
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.Create(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusSeeOther, rec.Code)
}

func TestGetCreate(t *testing.T) {
	mockUcase := new(mocks.TodoUsecase)
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/todo/create", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.TodoController{
		Usecase: mockUcase,
		Store:   store,
	}

	session, _ := store.Get(req, "session_id")
	session.Values["id"] = 1
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.Create(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestFindAll(t *testing.T) {
	mockUcase := new(mocks.TodoUsecase)
	userID := 1
	mockResponse := &dto.FindAllTodoResponse{
		Todos: []*domain.Todo{
			{
				ID:          1,
				Title:       "test",
				Description: "test",
				UserID:      1,
			},
			{
				ID:          2,
				Title:       "test2",
				Description: "test2",
				UserID:      1,
			},
		},
	}
	mockUcase.On("FindAll", mock.Anything, userID).Return(mockResponse, nil).Once()
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/todo", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.TodoController{
		Usecase: mockUcase,
		Store:   store,
	}

	session, _ := store.Get(req, "session_id")
	session.Values["id"] = userID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.FindAll(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestFindByID(t *testing.T) {
	mockUcase := new(mocks.TodoUsecase)
	todoID := 1
	mockUcase.On("FindByID", mock.Anything, todoID).Return(&dto.FindByIDTodoResponse{
		Todo: &domain.Todo{
			ID:          1,
			Title:       "test",
			Description: "test",
			UserID:      1,
		},
	}, nil).Once()
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/todo/"+strconv.Itoa(todoID), nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(todoID))

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.TodoController{
		Usecase: mockUcase,
		Store:   store,
	}

	userID := 1
	session, _ := store.Get(req, "session_id")
	session.Values["id"] = userID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.FindByID(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdate(t *testing.T) {
	mockUcase := new(mocks.TodoUsecase)
	request := dto.UpdateTodoRequest{
		ID:          1,
		Title:       "update",
		Description: "update",
	}
	j, err := json.Marshal(request)
	require.NoError(t, err)

	mockUcase.On("FindByID", mock.Anything, request.ID).Return(&dto.FindByIDTodoResponse{
		Todo: &domain.Todo{
			ID:          1,
			Title:       "test",
			Description: "test",
			UserID:      1,
		},
	}, nil).Once()
	mockUcase.On("Update", mock.Anything, request).Return(nil).Once()

	mockUcase.On("FindAll", mock.Anything, 1).Return(&dto.FindAllTodoResponse{
		Todos: []*domain.Todo{
			{
				ID:          1,
				Title:       "test",
				Description: "test",
				UserID:      1,
			},
			{
				ID:          2,
				Title:       "test2",
				Description: "test2",
				UserID:      1,
			},
		},
	}, nil).Once()
	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPut, "todo/"+strconv.Itoa(request.ID), strings.NewReader(string(j)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(request.ID))

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.TodoController{
		Usecase: mockUcase,
		Store:   store,
	}

	userID := 1
	session, _ := store.Get(req, "session_id")
	session.Values["id"] = userID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.Update(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestDelete(t *testing.T) {
	mockUcase := new(mocks.TodoUsecase)
	todoID := 1
	mockUcase.On("FindByID", mock.Anything, todoID).Return(&dto.FindByIDTodoResponse{
		Todo: &domain.Todo{
			ID:          1,
			Title:       "test",
			Description: "test",
			UserID:      1,
		},
	}, nil).Once()
	mockUcase.On("Delete", mock.Anything, todoID).Return(nil).Once()
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodDelete, "/todo/"+strconv.Itoa(todoID), nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(todoID))

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.TodoController{
		Usecase: mockUcase,
		Store:   store,
	}

	userID := 1
	session, _ := store.Get(req, "session_id")
	session.Values["id"] = userID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.Delete(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestIsDone(t *testing.T) {
	mockUcase := new(mocks.TodoUsecase)
	todoID := 1
	mockUcase.On("FindByID", mock.Anything, todoID).Return(&dto.FindByIDTodoResponse{
		Todo: &domain.Todo{
			ID:          1,
			Title:       "test",
			Description: "test",
			UserID:      1,
		},
	}, nil).Once()
	mockUcase.On("IsDone", mock.Anything, todoID).Return(nil).Once()
	mockUcase.On("FindAll", mock.Anything, 1).Return(&dto.FindAllTodoResponse{
		Todos: []*domain.Todo{
			{
				ID:          1,
				Title:       "test",
				Description: "test",
				UserID:      1,
			},
			{
				ID:          2,
				Title:       "test2",
				Description: "test2",
				UserID:      1,
			},
		},
	}, nil).Once()
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPut, "/todo/done/"+strconv.Itoa(todoID), nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(todoID))

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.TodoController{
		Usecase: mockUcase,
		Store:   store,
	}

	userID := 1
	session, _ := store.Get(req, "session_id")
	session.Values["id"] = userID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.IsDone(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestIsNotDone(t *testing.T) {
	mockUcase := new(mocks.TodoUsecase)
	todoID := 1
	mockUcase.On("FindByID", mock.Anything, todoID).Return(&dto.FindByIDTodoResponse{
		Todo: &domain.Todo{
			ID:          1,
			Title:       "test",
			Description: "test",
			UserID:      1,
		},
	}, nil).Once()
	mockUcase.On("IsNotDone", mock.Anything, todoID).Return(nil).Once()
	mockUcase.On("FindAll", mock.Anything, 1).Return(&dto.FindAllTodoResponse{
		Todos: []*domain.Todo{
			{
				ID:          1,
				Title:       "test",
				Description: "test",
				UserID:      1,
			},
			{
				ID:          2,
				Title:       "test2",
				Description: "test2",
				UserID:      1,
			},
		},
	}, nil).Once()
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPut, "/todo/notdone/"+strconv.Itoa(todoID), nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(todoID))

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.TodoController{
		Usecase: mockUcase,
		Store:   store,
	}

	userID := 1
	session, _ := store.Get(req, "session_id")
	session.Values["id"] = userID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.IsNotDone(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}
