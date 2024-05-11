package controller

import "github.com/gofiber/fiber/v2"

func parseAndHandleError(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	return nil
}

func handleError(c *fiber.Ctx, err error, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error": err.Error(),
	})
}
