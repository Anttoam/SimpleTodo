package controller

import (
	"context"
	"errors"
	"strconv"

	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/views/todo"
	"github.com/Anttoam/golang-htmx-todos/views/user"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type UserUsecase interface {
	SignUp(ctx context.Context, user dto.SignUpRequest) error
	Login(ctx context.Context, user dto.LoginRequest) (*dto.LoginResponse, error)
	FindByID(ctx context.Context, userID int) (*dto.FindByIDUserResponse, error)
	EditUser(ctx context.Context, user dto.UpdateUserRequest) error
	EditPassword(ctx context.Context, user dto.UpdatePasswordRequest) error
}

type UserController struct {
	uu    UserUsecase
	store *session.Store
}

func NewUserController(app *fiber.App, uu UserUsecase, store *session.Store) {
	uc := &UserController{uu: uu, store: store}

	api := app.Group("/user")
	api.Get("/signup", uc.SignUp)
	api.Post("/signup", uc.SignUp)
	api.Get("/login", uc.Login)
	api.Post("/login", uc.Login)
	api.Get("/logout", uc.Logout)
	api.Get("/:id", uc.EditUser)
	api.Put("/:id", uc.EditUser)
	api.Get("/password/:id", uc.EditPassword)
	api.Put("/password/:id", uc.EditPassword)
}

func (uc *UserController) SignUp(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		req := dto.SignUpRequest{
			Name:     c.FormValue("name"),
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}
		if err := parseAndHandleError(c, &req); err != nil {
			return err
		}

		ctx := c.Context()
		if err := uc.uu.SignUp(ctx, req); err != nil {
			return handleError(c, err, fiber.StatusInternalServerError)
		}

		return c.Redirect("/user/login")

	}

	component := templ.Handler(user.SignUp())
	handler := adaptor.HTTPHandler(component)
	return handler(c)
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		var req dto.LoginRequest
		if err := parseAndHandleError(c, &req); err != nil {
			return err
		}

		ctx := c.Context()
		res, err := uc.uu.Login(ctx, req)
		if err != nil {
			return handleError(c, err, fiber.StatusInternalServerError)
		}

		sess, _ := uc.store.Get(c)
		sess.Set("id", res.ID)
		if err := sess.Save(); err != nil {
			return handleError(c, err, fiber.StatusInternalServerError)
		}

		return c.Redirect("/todos")
	}

	component := templ.Handler(user.Login())
	handler := adaptor.HTTPHandler(component)
	return handler(c)
}

func (uc *UserController) Logout(c *fiber.Ctx) error {
	sess, _ := uc.store.Get(c)
	if err := sess.Destroy(); err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}
	return c.Redirect("/user/login")
}

func (uc *UserController) EditUser(c *fiber.Ctx) error {
	var userID int
	sess, _ := uc.store.Get(c)
	id := sess.Get("id")
	if id == nil {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}
	userID = id.(int)

	var fetch *dto.FindByIDUserResponse

	if c.Method() == fiber.MethodPut {
		var req dto.UpdateUserRequest
		if err := parseAndHandleError(c, &req); err != nil {
			return err
		}

		sess, _ := uc.store.Get(c)
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
		fetch, err := uc.uu.FindByID(ctx, req.ID)
		if err != nil {
			return handleError(c, err, fiber.StatusInternalServerError)
		}

		if fetch.User.ID != userID {
			return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
		}

		if err := uc.uu.EditUser(ctx, req); err != nil {
			return handleError(c, err, fiber.StatusInternalServerError)
		}

	}

	ctx := c.Context()
	fetch, err := uc.uu.FindByID(ctx, userID)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}

	component := templ.Handler(todo.EditUser(strconv.Itoa(userID), fetch.User.Name, fetch.User.Email))
	handler := adaptor.HTTPHandler(component)
	return handler(c)
}

func (uc *UserController) EditPassword(c *fiber.Ctx) error {
	var userID int
	sess, _ := uc.store.Get(c)
	id := sess.Get("id")
	if id == nil {
		return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
	}
	userID = id.(int)

	if c.Method() == fiber.MethodPut {
		var req dto.UpdatePasswordRequest
		if err := parseAndHandleError(c, &req); err != nil {
			return err
		}

		sess, _ := uc.store.Get(c)
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
		fetch, err := uc.uu.FindByID(ctx, req.ID)
		if err != nil {
			return handleError(c, err, fiber.StatusInternalServerError)
		}

		req.Password = c.FormValue("password")
		req.NewPassword = c.FormValue("new_password")

		if fetch.User.ID != userID {
			return handleError(c, errors.New("Unauthorized"), fiber.StatusUnauthorized)
		}

		if err := uc.uu.EditPassword(ctx, req); err != nil {
			return handleError(c, err, fiber.StatusInternalServerError)
		}

	}

	component := templ.Handler(todo.EditPassword(strconv.Itoa(userID)))
	handler := adaptor.HTTPHandler(component)
	return handler(c)
}
