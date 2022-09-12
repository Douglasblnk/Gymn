package equipmentService

import (
	"gymn/internal/dto"
	equipmentRepository "gymn/internal/repository/equipment"
	"gymn/internal/utils"
)

func GetEquipment(id string) (*dto.EquipmentDTO, *utils.Error) {
	equipment, err := equipmentRepository.GetEquipmentByUID(id)

	if err != nil {
		return nil, err
	}

	return &dto.EquipmentDTO{
		Name: equipment.Name,
	}, nil
}
