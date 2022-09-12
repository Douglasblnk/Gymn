package equipmentService

import (
	"gymn/internal/dto"
	"gymn/internal/models"
	equipmentRepository "gymn/internal/repository/equipment"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func CreateEquipment(userID int, data *schemas.RegisterEquipment) (*dto.EquipmentDTO, *utils.Error) {
	_, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return nil, err
	}

	equipment := &models.Equipment{
		Name: data.Name,
	}

	if err := equipmentRepository.CreateEquipment(equipment); err != nil {
		return nil, err
	}

	return &dto.EquipmentDTO{
		Name: equipment.Name,
	}, nil
}
