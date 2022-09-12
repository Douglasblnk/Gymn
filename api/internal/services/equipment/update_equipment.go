package equipmentService

import (
	"gymn/internal/dto"
	equipmentRepository "gymn/internal/repository/equipment"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func UpdateEquipment(userID int, id string, data *schemas.UpdateEquipment) (*dto.EquipmentDTO, *utils.Error) {
	_, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return nil, err
	}

	equipment, err := equipmentRepository.GetEquipmentByUID(id)

	if err != nil {
		return nil, err
	}

	if data.Name != nil {
		equipment.Name = *data.Name
	}

	if err = equipmentRepository.UpdateEquipment(equipment); err != nil {
		return nil, err
	}

	return &dto.EquipmentDTO{
		Name: equipment.Name,
	}, nil
}
