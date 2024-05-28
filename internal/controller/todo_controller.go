package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/views/todo"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v9"
)

type TodoUsecase interface {
	Create(ctx context.Context, req dto.CreateTodoRequest) error
	FindAll(ctx context.Context, userID int) (*dto.FindAllTodoResponse, error)
	FindByID(ctx context.Context, todoID int) (*dto.FindByIDTodoResponse, error)
	Update(ctx context.Context, req dto.UpdateTodoRequest) error
	Delete(ctx context.Context, todoID int) error
	IsDone(ctx context.Context, todoID int) error
	IsNotDone(ctx context.Context, todoID int) error
}

type TodoController struct {
	tu    TodoUsecase
	store *redisstore.RedisStore
}

func NewTodoController(e *echo.Echo, tu TodoUsecase, store *redisstore.RedisStore) {
	t := &TodoController{tu: tu, store: store}

	api := e.Group("/todo")
	api.GET("/create", t.Create)
	api.POST("/create", t.Create)
	api.GET("/", t.FindAll)
	api.GET("/:id", t.FindByID)
	api.PUT("/:id", t.Update)
	api.DELETE("/:id", t.Delete)
	api.PUT("/done/:id", t.IsDone)
	api.PUT("/notdone/:id", t.IsNotDone)
}

func (t *TodoController) Create(c echo.Context) error {
	sess, _ := t.store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}

	if c.Request().Method == http.MethodPost {
		req := dto.CreateTodoRequest{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
		}
		if err := c.Bind(&req); err != nil {
			return err
		}

		req.UserID = id.(int)

		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		ctx := c.Request().Context()
		if err := t.tu.Create(ctx, req); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusSeeOther, "/todo/")
	}

	component := todo.Create()
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) FindAll(c echo.Context) error {
	sess, _ := t.store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}

	ctx := c.Request().Context()
	userID := id.(int)
	res, err := t.tu.FindAll(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	component := todo.Page(*res, strconv.Itoa(userID))
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) FindByID(c echo.Context) error {
	sess, _ := t.store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	todoID := idP

	ctx := c.Request().Context()
	res, err := t.tu.FindByID(ctx, todoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if res.Todo.UserID != id.(int) {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}

	component := todo.EditForm(strconv.Itoa(res.Todo.ID), res.Todo.Title, res.Todo.Description)
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) Update(c echo.Context) error {
	req := dto.UpdateTodoRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}
	if err := c.Bind(&req); err != nil {
		return err
	}

	sess, _ := t.store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}
	userID := id.(int)

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	req.ID = idP

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	fetch, err := t.tu.FindByID(ctx, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if fetch.Todo.UserID != userID {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}
	if err := t.tu.Update(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := t.tu.FindAll(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	component := todo.List(*res)
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) Delete(c echo.Context) error {
	sess, _ := t.store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	ctx := c.Request().Context()
	res, err := t.tu.FindByID(ctx, idP)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if res.Todo.UserID != id.(int) {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}

	if err := t.tu.Delete(ctx, res.Todo.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	component := todo.Delete(strconv.Itoa(res.Todo.ID))
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) IsDone(c echo.Context) error {
	return t.changeStatus(c, t.tu.IsDone)
}

func (t *TodoController) IsNotDone(c echo.Context) error {
	return t.changeStatus(c, t.tu.IsNotDone)
}
