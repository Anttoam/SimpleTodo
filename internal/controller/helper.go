package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Anttoam/golang-htmx-todos/views/todo"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func (t *TodoController) changeStatus(c echo.Context, changeFunc func(ctx context.Context, todoID int) error) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	sess, _ := t.store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}
	userID := id.(int)

	ctx := c.Request().Context()
	fetch, err := t.tu.FindByID(ctx, idP)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if fetch.Todo.UserID != userID {
		return c.JSON(http.StatusUnauthorized, errors.New("Unauthorized").Error())
	}

	if err := changeFunc(ctx, idP); err != nil {
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
