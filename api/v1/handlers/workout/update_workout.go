package handlers

import (
	workoutService "gymn/internal/services/workout"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func UpdateWorkout(c *fiber.Ctx) error {
	workoutSchema := new(schemas.UpdateWorkout)
	workoutId := c.Params("id")

	if err := utils.GetBody(workoutSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userID := c.Locals("userID").(int)

	response, err := workoutService.UpdateWorkout(userID, workoutId, workoutSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)

}
