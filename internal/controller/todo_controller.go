package controller

import (
	"context"
	"errors"

	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type TodoUsecase interface {
	Create(ctx context.Context, req dto.CreateTodoRequest) (*dto.CreateTodoResponse, error)
}

type TodoController struct {
	tu    TodoUsecase
	store *session.Store
}

func NewTodoController(app *fiber.App, tu TodoUsecase, store *session.Store) {
	todo := &TodoController{tu: tu, store: store}

	app.Post("/create", todo.Create)
}

func (t *TodoController) Create(c *fiber.Ctx) error {
	sess, _ := t.store.Get(c)
	id := sess.Get("id")
	if id == nil {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}

	var req dto.CreateTodoRequest
	if err := parseAndHandleError(c, &req); err != nil {
		return err
	}

	req.UserID = id.(int)

	ctx := c.Context()
	res, err := t.tu.Create(ctx, req)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
