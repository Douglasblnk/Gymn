package handlers

import (
	workoutService "gymn/internal/services/workout"

	"github.com/gofiber/fiber/v2"
)

func DeleteWorkout(c *fiber.Ctx) error {
	workoutID := c.Params("id")
	userID := c.Locals("userID").(int)

	response, err := workoutService.DeleteWorkout(userID, workoutID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).SendString(response)
}
