package handlers

import (
	studentService "gymn/internal/services/student"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateStudent(c *fiber.Ctx) error {
	studentSchema := new(schemas.RegisterStudent)

	if err := utils.GetBody(studentSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userID := c.Locals("userID").(int)

	response, err := studentService.CreateStudent(userID, studentSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
