package handlers

import (
	trainingSheetService "gymn/internal/services/training-sheet"

	"github.com/gofiber/fiber/v2"
)

func GetStudentTrainingSheets(c *fiber.Ctx) error {
	studentID := c.Params("id")
	userID := c.Locals("userID").(int)

	response, err := trainingSheetService.GetStudentTrainingSheets(userID, studentID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
