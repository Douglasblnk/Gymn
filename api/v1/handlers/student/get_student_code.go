package handlers

import (
	studentService "gymn/internal/services/student"

	"github.com/gofiber/fiber/v2"
)

func GetStudentCode(c *fiber.Ctx) error {
	studentUID := c.Params("id")
	userID := c.Locals("userID").(int)

	response, err := studentService.GetStudentCode(userID, studentUID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
