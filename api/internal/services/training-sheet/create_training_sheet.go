package trainingSheetService

import (
	"gymn/internal/dto"
	"gymn/internal/models"
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func CreateTrainingSheet(userID int, data *schemas.RegisterTrainingSheet) (*dto.TrainingSheetDTO, *utils.Error) {
	user, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return nil, err
	}

	trainingSheet := &models.TrainingSheet{
		Name: data.Name,
	}

	if err := trainingSheetRepository.CreateTrainingSheet(user, trainingSheet); err != nil {
		return nil, err
	}

	return &dto.TrainingSheetDTO{
		Name:   trainingSheet.Name,
		Active: true,
	}, nil
}
