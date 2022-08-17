package handlers

import (
	studentService "gymn/internal/services/student"

	"github.com/gofiber/fiber/v2"
)

func GetStudents(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	response, err := studentService.GetStudents(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
