package controller

import (
	"context"

	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/gofiber/fiber/v2"
)

type UserUsecase interface {
	SignUp(ctx context.Context, user dto.SignUpRequest) error
	Login(ctx context.Context, user dto.LoginRequest) error
}

type UserController struct {
	uu UserUsecase
}

func NewUserController(app *fiber.App, uu UserUsecase) {
	user := &UserController{uu: uu}

	app.Post("/signup", user.SignUp)
	app.Post("/login", user.Login)
}

func (u *UserController) SignUp(c *fiber.Ctx) error {
	var user dto.SignUpRequest
	if err := parseAndHandleError(c, &user); err != nil {
		return err
	}

	ctx := c.Context()
	if err := u.uu.SignUp(ctx, user); err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (u *UserController) Login(c *fiber.Ctx) error {
	var user dto.LoginRequest
	if err := parseAndHandleError(c, &user); err != nil {
		return err
	}

	ctx := c.Context()
	if err := u.uu.Login(ctx, user); err != nil {
		return handleError(c, err, fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged in successfully",
	})
}
