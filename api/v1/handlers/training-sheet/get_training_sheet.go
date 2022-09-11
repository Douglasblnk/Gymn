package handlers

import (
	trainingSheetService "gymn/internal/services/training-sheet"

	"github.com/gofiber/fiber/v2"
)

func GetTrainingSheet(c *fiber.Ctx) error {
	trainingSheetID := c.Params("id")
	userID := c.Locals("userID").(int)

	response, err := trainingSheetService.GetTrainingSheet(userID, trainingSheetID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
