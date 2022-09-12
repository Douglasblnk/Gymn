package equipmentRepository

import (
	"gymn/internal/database"
	"gymn/internal/exceptions"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetEquipmentByUID(uid string) (*models.Equipment, *utils.Error) {
	equipment := &models.Equipment{}

	if err := database.DB.Where("uid = ?", uid).First(equipment).Error; err != nil {
		return nil, utils.Throw(exceptions.ErrEquipmentNotFound, 404)
	}

	return equipment, nil
}
