package handlers

import (
	userService "gymn/internal/services/user"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	userSchema := new(schemas.User)

	if err := utils.GetBody(userSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	response, err := userService.RegisterUser(userSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
