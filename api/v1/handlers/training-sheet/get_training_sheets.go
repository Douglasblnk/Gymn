package handlers

import (
	trainingSheetService "gymn/internal/services/training-sheet"

	"github.com/gofiber/fiber/v2"
)

func GetTrainingSheets(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	response, err := trainingSheetService.GetTrainingSheets(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
