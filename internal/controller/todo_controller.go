package controller

import (
	"context"
	"errors"
	"strconv"

	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/views/todo"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/session"
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
	store *session.Store
}

func NewTodoController(app *fiber.App, tu TodoUsecase, store *session.Store) {
	t := &TodoController{tu: tu, store: store}

	app.Get("/create", t.Create)
	app.Post("/create", t.Create)
	app.Get("/todos", t.FindAll)
	app.Get("/:id", t.FindByID)
	app.Put("/:id", t.Update)
	app.Delete("/:id", t.Delete)
	app.Put("/done/:id", t.IsDone)
	app.Put("/notdone/:id", t.IsNotDone)
}

func (t *TodoController) Create(c *fiber.Ctx) error {
	sess, _ := t.store.Get(c)
	id := sess.Get("id")
	if id == nil {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}

	if c.Method() == fiber.MethodPost {
		var req dto.CreateTodoRequest
		if err := parseAndHandleError(c, &req); err != nil {
			return err
		}

		req.UserID = id.(int)

		ctx := c.Context()
		if err := t.tu.Create(ctx, req); err != nil {
			return handleError(c, err, fiber.StatusInternalServerError)
		}

		return c.Redirect("/todos")
	}

	component := todo.Create()
	handler := adaptor.HTTPHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) FindAll(c *fiber.Ctx) error {
	sess, _ := t.store.Get(c)
	id := sess.Get("id")
	if id == nil {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}

	ctx := c.Context()
	userID := id.(int)
	res, err := t.tu.FindAll(ctx, userID)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}

	component := todo.Page(*res)
	handler := adaptor.HTTPHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) FindByID(c *fiber.Ctx) error {
	sess, _ := t.store.Get(c)
	id := sess.Get("id")
	if id == nil {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}

	idP, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return handleError(c, err, fiber.StatusNotFound)
	}

	todoID := idP

	ctx := c.Context()
	res, err := t.tu.FindByID(ctx, todoID)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}
	if res.Todo.UserID != id.(int) {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}

	component := todo.EditForm(strconv.Itoa(res.Todo.ID), res.Todo.Title, res.Todo.Description)
	handler := adaptor.HTTPHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) Update(c *fiber.Ctx) error {
	var req dto.UpdateTodoRequest
	if err := parseAndHandleError(c, &req); err != nil {
		return err
	}

	sess, _ := t.store.Get(c)
	id := sess.Get("id")
	if id == nil {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}
	userID := id.(int)

	idP, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return handleError(c, err, fiber.StatusNotFound)
	}

	req.ID = idP

	ctx := c.Context()
	fetch, err := t.tu.FindByID(ctx, req.ID)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}

	if fetch.Todo.UserID != userID {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}
	if err := t.tu.Update(ctx, req); err != nil {
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

func (t *TodoController) Delete(c *fiber.Ctx) error {
	sess, _ := t.store.Get(c)
	id := sess.Get("id")
	if id == nil {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}

	idP, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return handleError(c, err, fiber.StatusNotFound)
	}

	ctx := c.Context()
	res, err := t.tu.FindByID(ctx, idP)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}

	if res.Todo.UserID != id.(int) {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}

	if err := t.tu.Delete(ctx, res.Todo.ID); err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}

	component := todo.Delete(strconv.Itoa(res.Todo.ID))
	handler := adaptor.HTTPHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) IsDone(c *fiber.Ctx) error {
	return t.changeStatus(c, t.tu.IsDone)
}

func (t *TodoController) IsNotDone(c *fiber.Ctx) error {
	return t.changeStatus(c, t.tu.IsNotDone)
}
