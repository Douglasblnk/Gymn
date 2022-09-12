package handlers

import (
	equipmentService "gymn/internal/services/equipment"

	"github.com/gofiber/fiber/v2"
)

func DeleteEquipment(c *fiber.Ctx) error {
	equipmentID := c.Params("id")
	userID := c.Locals("userID").(int)

	response, err := equipmentService.DeleteEquipment(userID, equipmentID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).SendString(response)
}
