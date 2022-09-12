package handlers

import (
	equipmentService "gymn/internal/services/equipment"
	"gymn/v1/schemas"
	"gymn/v1/utils"

	"github.com/gofiber/fiber/v2"
)

func UpdateEquipment(c *fiber.Ctx) error {
	equipmentSchema := new(schemas.UpdateEquipment)
	equipmentId := c.Params("id")

	if err := utils.GetBody(equipmentSchema, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userID := c.Locals("userID").(int)

	response, err := equipmentService.UpdateEquipment(userID, equipmentId, equipmentSchema)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
