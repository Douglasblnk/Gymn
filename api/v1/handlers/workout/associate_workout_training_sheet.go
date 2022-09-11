package handlers

import (
	workoutService "gymn/internal/services/workout"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func AssociateWorkoutTrainingSheet(c *fiber.Ctx) error {
	workoutTrainingSheetSchema := new(schemas.WorkoutTrainingSheet)

	if err := utils.GetBody(workoutTrainingSheetSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userID := c.Locals("userID").(int)

	response, err := workoutService.AssociateWorkoutTrainingSheet(userID, workoutTrainingSheetSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
