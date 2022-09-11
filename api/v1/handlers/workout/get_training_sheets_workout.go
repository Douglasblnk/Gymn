package handlers

import (
	workoutService "gymn/internal/services/workout"

	"github.com/gofiber/fiber/v2"
)

func GetTrainingSheetWorkouts(c *fiber.Ctx) error {
	trainingSheetID := c.Params("id")
	userID := c.Locals("userID").(int)

	response, err := workoutService.GetTrainingSheetWorkouts(userID, trainingSheetID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
