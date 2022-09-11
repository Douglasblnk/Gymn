package trainingSheetService

import (
	"gymn/internal/dto"
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func UpdateTrainingSheet(userID int, id string, data *schemas.UpdateTrainingSheet) (*dto.TrainingSheetDTO, *utils.Error) {
	_, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return nil, err
	}

	trainingSheet, err := trainingSheetRepository.GetTrainingSheetByUID(userID, id)

	if err != nil {
		return nil, err
	}

	if data.Name != nil {
		trainingSheet.Name = *data.Name
	}

	if data.Active != nil {
		trainingSheet.Active = *data.Active
	}

	if err = trainingSheetRepository.UpdateTrainingSheet(trainingSheet); err != nil {
		return nil, err
	}

	return &dto.TrainingSheetDTO{
		Name:   trainingSheet.Name,
		Active: trainingSheet.Active,
	}, nil
}
