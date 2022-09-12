package equipmentRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func DeleteEquipment(equipment *models.Equipment) *utils.Error {
	if err := database.DB.Delete(equipment).Error; err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
