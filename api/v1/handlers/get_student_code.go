package handlers

import (
	studentService "gymn/internal/services/student"

	"github.com/gofiber/fiber/v2"
)

func GetStudentCode(c *fiber.Ctx) error {
	studentUID := c.Params("id")

	response, err := studentService.GetStudentCode(studentUID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)

}
