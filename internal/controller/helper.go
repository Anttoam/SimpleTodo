package controller

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/Anttoam/golang-htmx-todos/views/todo"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func parseAndHandleError(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	return nil
}

func handleError(c *fiber.Ctx, err error, statusCode int) error {
	log.Println(err)
	return c.Status(statusCode).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func (t *TodoController) changeStatus(c *fiber.Ctx, changeFunc func(ctx context.Context, todoID int) error) error {
	idP, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return handleError(c, err, fiber.StatusNotFound)
	}

	sess, _ := t.store.Get(c)
	id := sess.Get("id")
	if id == nil {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}
	userID := id.(int)

	ctx := c.Context()
	fetch, err := t.tu.FindByID(ctx, idP)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}

	if fetch.Todo.UserID != userID {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}

	if err := changeFunc(ctx, idP); err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}

	res, err := t.tu.FindAll(ctx, userID)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}

	component := todo.List(*res)
	handler := adaptor.HTTPHandler(templ.Handler(component))
	return handler(c)
}
