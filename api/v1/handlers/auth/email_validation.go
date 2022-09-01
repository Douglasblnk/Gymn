package handlers

import (
	authService "gymn/internal/services/auth"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

type Email struct {
	Email string `json:"email" validate:"required"`
}

func EmailValidation(c *fiber.Ctx) error {
	email := new(Email)

	if err := utils.GetBody(email, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	response, err := authService.ValidateEmail(email.Email)

	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
