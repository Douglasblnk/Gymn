package handlers

import (
	studentService "gymn/internal/services/student"

	"github.com/gofiber/fiber/v2"
)

func GetStudent(c *fiber.Ctx) error {
	studentID := c.Params("id")
	userID := c.Locals("userID").(int)

	response, err := studentService.GetStudent(userID, studentID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
