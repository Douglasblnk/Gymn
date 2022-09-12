package handlers

import (
	equipmentService "gymn/internal/services/equipment"

	"github.com/gofiber/fiber/v2"
)

func GetEquipment(c *fiber.Ctx) error {
	equipmentID := c.Params("id")

	response, err := equipmentService.GetEquipment(equipmentID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
