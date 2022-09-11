package handlers

import (
	studentService "gymn/internal/services/student"

	"github.com/gofiber/fiber/v2"
)

func GetAllStudents(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	response, err := studentService.GetAllStudents(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
