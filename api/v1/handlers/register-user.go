package handlers

import (
	"fmt"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	userSchema := new(schemas.User)

	if err := utils.GetBody(userSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	fmt.Println("userSchema", userSchema)
	// response, err := authService.SignIn(loginSchema)

	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(err)
	// }

	// return c.Status(fiber.StatusCreated).JSON(response)

	return nil
}
