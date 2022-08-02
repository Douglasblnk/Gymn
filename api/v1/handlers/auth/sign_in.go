package handlers

import (
	authService "gymn/internal/services/auth"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func SignInHandler(c *fiber.Ctx) error {
	loginSchema := new(schemas.Login)

	if err := utils.GetBody(loginSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	response, err := authService.SignIn(loginSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
