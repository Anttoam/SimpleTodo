package controller

import (
	"context"

	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/views/user"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type UserUsecase interface {
	SignUp(ctx context.Context, user dto.SignUpRequest) error
	Login(ctx context.Context, user dto.LoginRequest) (*dto.LoginResponse, error)
}

type UserController struct {
	uu    UserUsecase
	store *session.Store
}

func NewUserController(app *fiber.App, uu UserUsecase, store *session.Store) {
	uc := &UserController{uu: uu, store: store}

	app.Get("/signup", uc.SignUp)
	app.Post("/signup", uc.SignUp)
	app.Get("/login", uc.Login)
	app.Post("/login", uc.Login)
	app.Get("/logout", uc.Logout)
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

		return c.Redirect("/login")

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
	return c.Redirect("/login")
}
