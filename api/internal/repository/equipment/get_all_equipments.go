package equipmentRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetAllEquipments() ([]*models.Equipment, *utils.Error) {
	var equipments []*models.Equipment

	query := database.DB.Order("created_at desc")

	if err := query.Find(&equipments).Error; err != nil {
		return nil, utils.Throw(err.Error(), 400)
	}

	return equipments, nil
}
