package handlers

import (
	studentService "gymn/internal/services/student"

	"github.com/gofiber/fiber/v2"
)

func DeleteStudent(c *fiber.Ctx) error {
	studentID := c.Params("id")
	userID := c.Locals("userID").(int)

	response, err := studentService.DeleteStudent(userID, studentID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).SendString(response)
}
