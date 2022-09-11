package handlers

import (
	trainingSheetService "gymn/internal/services/training-sheet"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func AssociateTrainingSheetStudent(c *fiber.Ctx) error {
	studentTrainingSheetSchema := new(schemas.StudentTrainingSheet)

	if err := utils.GetBody(studentTrainingSheetSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userID := c.Locals("userID").(int)

	response, err := trainingSheetService.AssociateTrainingSheetStudent(userID, studentTrainingSheetSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
