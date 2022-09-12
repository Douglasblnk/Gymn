package equipmentService

import (
	"gymn/internal/dto"
	equipmentRepository "gymn/internal/repository/equipment"
	"gymn/internal/utils"
)

func GetAllEquipments() ([]*dto.EquipmentDTO, *utils.Error) {
	equipmentsModel, err := equipmentRepository.GetAllEquipments()

	if err != nil {
		return nil, err
	}

	var equipments []*dto.EquipmentDTO

	for _, equipment := range equipmentsModel {
		equipmentDTO := &dto.EquipmentDTO{
			Name: equipment.Name,
		}

		equipments = append(equipments, equipmentDTO)
	}

	return equipments, nil
}
