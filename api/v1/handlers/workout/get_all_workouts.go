package handlers

import (
	workoutService "gymn/internal/services/workout"

	"github.com/gofiber/fiber/v2"
)

func GetAllWorkouts(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	response, err := workoutService.GetAllWorkouts(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
