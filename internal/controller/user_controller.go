package controller

import (
	"context"

	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type UserUsecase interface {
	SignUp(ctx context.Context, user dto.SignUpRequest) (*dto.SignUpResponse, error)
	Login(ctx context.Context, user dto.LoginRequest) (*dto.LoginResponse, error)
}

type UserController struct {
	uu    UserUsecase
	store *session.Store
}

func NewUserController(app *fiber.App, uu UserUsecase, store *session.Store) {
	user := &UserController{uu: uu, store: store}

	app.Post("/signup", user.SignUp)
	app.Post("/login", user.Login)
}

func (uc *UserController) SignUp(c *fiber.Ctx) error {
	var req dto.SignUpRequest
	if err := parseAndHandleError(c, &req); err != nil {
		return err
	}

	ctx := c.Context()
	res, err := uc.uu.SignUp(ctx, req)
	if err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func (uc *UserController) Login(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusOK).JSON(res)
}
