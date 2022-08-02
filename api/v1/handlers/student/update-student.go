package handlers

import (
	studentService "gymn/internal/services/student"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func UpdateStudent(c *fiber.Ctx) error {
	studentSchema := new(schemas.UpdateStudent)
	studentId := c.Params("id")

	if err := utils.GetBody(studentSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	response, err := studentService.UpdateStudent(studentId, studentSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
