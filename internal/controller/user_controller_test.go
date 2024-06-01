package controller_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/Anttoam/SimpleTodo/domain"
	"github.com/Anttoam/SimpleTodo/dto"
	"github.com/Anttoam/SimpleTodo/internal/controller"
	"github.com/Anttoam/SimpleTodo/internal/controller/mocks"
	"github.com/Anttoam/SimpleTodo/pkg/validation"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	mockUcase := new(mocks.UserUsecase)
	request := dto.SignUpRequest{
		Name:     "test",
		Email:    "test@example.com",
		Password: "password",
	}
	j, err := json.Marshal(request)
	require.NoError(t, err)

	mockUcase.On("SignUp", mock.Anything, request).Return(nil).Once()

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "user/signup", strings.NewReader(string(j)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	uc := controller.UserController{
		Usecase: mockUcase,
		Store:   nil,
	}

	err = uc.SignUp(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusSeeOther, rec.Code)
}

func TestSignUpValidationError(t *testing.T) {
	mockUcase := new(mocks.UserUsecase)
	request := dto.SignUpRequest{
		Name:     "test",
		Email:    "test@example.com",
		Password: "pass",
	}
	j, err := json.Marshal(request)
	require.NoError(t, err)

	mockUcase.On("SignUp", mock.Anything, request).Return(nil).Once()
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "user/signup", strings.NewReader(string(j)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	uc := controller.UserController{
		Usecase: mockUcase,
		Store:   nil,
	}

	err = uc.SignUp(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetSignUp(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "user/signup", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	uc := controller.UserController{
		Usecase: nil,
		Store:   nil,
	}

	err = uc.SignUp(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestLogin(t *testing.T) {
	mockUcase := new(mocks.UserUsecase)
	request := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}
	j, err := json.Marshal(request)
	require.NoError(t, err)

	mockUcase.On("Login", mock.Anything, request).Return(&dto.LoginResponse{
		ID:    1,
		Email: "test@example.com",
	}, nil).Once()

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "user/login", strings.NewReader(string(j)))
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

	uc := controller.UserController{
		Usecase: mockUcase,
		Store:   store,
	}

	err = uc.Login(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusSeeOther, rec.Code)
}

func TestGetLogin(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "user/login", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	uc := controller.UserController{
		Usecase: nil,
		Store:   nil,
	}

	err = uc.Login(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestLogout(t *testing.T) {
	mockUcase := new(mocks.UserUsecase)

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, "user/logout", nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.UserController{
		Usecase: mockUcase,
		Store:   store,
	}

	err = uc.Logout(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusFound, rec.Code)
}

func TestEditUser(t *testing.T) {
	mockUcase := new(mocks.UserUsecase)
	request := dto.UpdateUserRequest{
		ID:    1,
		Name:  "updated",
		Email: "updated@example.com",
	}
	j, err := json.Marshal(request)
	require.NoError(t, err)

	mockUcase.On("FindUserByID", mock.Anything, request.ID).Return(&dto.FindByIDUserResponse{
		User: &domain.User{
			ID:    1,
			Name:  "test",
			Email: "test@example.com",
		},
	}, nil).Once()
	mockUcase.On("EditUser", mock.Anything, request).Return(nil).Once()

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPut, "user/"+strconv.Itoa(request.ID), strings.NewReader(string(j)))
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

	uc := controller.UserController{
		Usecase: mockUcase,
		Store:   store,
	}

	session, _ := store.Get(req, "session_id")
	session.Values["id"] = request.ID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.EditUser(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusSeeOther, rec.Code)
}

func TestGetEditUser(t *testing.T) {
	mockUcase := new(mocks.UserUsecase)
	userID := 1

	mockUcase.On("FindUserByID", mock.Anything, userID).Return(&dto.FindByIDUserResponse{
		User: &domain.User{
			ID:    1,
			Name:  "test",
			Email: "test@example.com",
		},
	}, nil).Once()

	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "user/"+strconv.Itoa(userID), nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(userID))

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.UserController{
		Usecase: mockUcase,
		Store:   store,
	}

	session, _ := store.Get(req, "session_id")
	session.Values["id"] = userID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.EditUser(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestEditPassword(t *testing.T) {
	mockUcase := new(mocks.UserUsecase)
	request := dto.UpdatePasswordRequest{
		ID:          1,
		Password:    "password",
		NewPassword: "new_password",
	}
	j, err := json.Marshal(request)
	require.NoError(t, err)

	mockUcase.On("FindUserByID", mock.Anything, request.ID).Return(&dto.FindByIDUserResponse{
		User: &domain.User{
			ID:       1,
			Name:     "test",
			Email:    "test@example.com",
			Password: "password",
		},
	}, nil).Once()
	mockUcase.On("EditPassword", mock.Anything, request).Return(nil).Once()

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPut, "user/password/"+strconv.Itoa(request.ID), strings.NewReader(string(j)))
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

	uc := controller.UserController{
		Usecase: mockUcase,
		Store:   store,
	}

	session, _ := store.Get(req, "session_id")
	session.Values["id"] = request.ID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.EditPassword(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusSeeOther, rec.Code)
}

func TestGetEditPassword(t *testing.T) {
	mockUcase := new(mocks.UserUsecase)
	userID := 1
	mockUcase.On("FindUserByID", mock.Anything, userID).Return(&dto.FindByIDUserResponse{
		User: &domain.User{
			ID:    1,
			Name:  "test",
			Email: "test@example.com",
		},
	}, nil).Once()

	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "user/password/"+strconv.Itoa(userID), nil)
	require.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(userID))

	s := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	require.NoError(t, err)

	uc := controller.UserController{
		Usecase: mockUcase,
		Store:   store,
	}

	session, _ := store.Get(req, "session_id")
	session.Values["id"] = userID
	err = session.Save(req, rec)
	require.NoError(t, err)

	err = uc.EditPassword(c)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}
