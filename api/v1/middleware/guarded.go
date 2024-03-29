package middleware

import (
	authService "gymn/internal/services/auth"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func JWTAuth(c *fiber.Ctx) error {
	authorization := utils.GetAuthorization(c)

	session, err := authService.ValidateAccessToken(authorization)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	c.Locals("userID", session.UserID)

	return c.Next()
}
