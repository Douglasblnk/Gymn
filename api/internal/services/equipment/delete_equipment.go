package equipmentService

import (
	equipmentRepository "gymn/internal/repository/equipment"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
)

func DeleteEquipment(userID int, studentID string) (string, *utils.Error) {
	_, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return "", err
	}

	equipment, err := equipmentRepository.GetEquipmentByUID(studentID)

	if err != nil {
		return "", err
	}

	if err := equipmentRepository.DeleteEquipment(equipment); err != nil {
		return "", err
	}

	return "deleted", nil
}
