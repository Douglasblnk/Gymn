package handlers

import (
	equipmentService "gymn/internal/services/equipment"

	"github.com/gofiber/fiber/v2"
)

func GetAllEquipments(c *fiber.Ctx) error {
	response, err := equipmentService.GetAllEquipments()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
