package handlers

import (
	trainingSheetService "gymn/internal/services/training-sheet"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateTrainingSheet(c *fiber.Ctx) error {
	trainingSheetSchema := new(schemas.RegisterTrainingSheet)

	if err := utils.GetBody(trainingSheetSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userID := c.Locals("userID").(int)

	response, err := trainingSheetService.CreateTrainingSheet(userID, trainingSheetSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)

}
