package equipmentRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func UpdateEquipment(equipment *models.Equipment) *utils.Error {
	if err := database.DB.Save(equipment).Error; err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
