package handlers

import (
	trainingSheetService "gymn/internal/services/training-sheet"

	"github.com/gofiber/fiber/v2"
)

func GetAllTrainingSheets(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	response, err := trainingSheetService.GetAllTrainingSheets(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
