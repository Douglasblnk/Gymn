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

	cookie := &fiber.Cookie{
		Name:     "AccessToken",
		Value:    response.AccessToken.Token,
		Expires:  response.AccessToken.Expiration,
		HTTPOnly: true,
		Secure:   utils.IsProductionEnv(),
		SameSite: (func() string {
			if utils.IsProductionEnv() {
				return "false"
			}

			return "none"
		})(),
	}

	c.Cookie(cookie)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
