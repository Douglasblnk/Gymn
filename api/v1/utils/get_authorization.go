package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAuthorization(c *fiber.Ctx) string {
	authorization := c.Get("authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		authorization = strings.ReplaceAll(authorization, "Bearer ", "")
	}

	return authorization
}
