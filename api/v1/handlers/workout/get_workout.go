package handlers

import (
	workoutService "gymn/internal/services/workout"

	"github.com/gofiber/fiber/v2"
)

func GetWorkout(c *fiber.Ctx) error {
	workoutID := c.Params("id")
	userID := c.Locals("userID").(int)

	response, err := workoutService.GetWorkout(userID, workoutID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
