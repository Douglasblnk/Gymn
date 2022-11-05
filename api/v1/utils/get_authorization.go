package utils

import (
	"github.com/gofiber/fiber/v2"
)

func GetAuthorization(c *fiber.Ctx) string {
	cookie := c.Cookies("AccessToken")

	return cookie
}
