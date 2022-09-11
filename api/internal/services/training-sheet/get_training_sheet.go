package trainingSheetService

import (
	"gymn/internal/dto"
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	"gymn/internal/utils"
)

func GetTrainingSheet(userID int, id string) (*dto.TrainingSheetDTO, *utils.Error) {
	trainingSheets, err := trainingSheetRepository.GetTrainingSheetByUID(userID, id)

	if err != nil {
		return nil, err
	}

	return &dto.TrainingSheetDTO{
		Name:   trainingSheets.Name,
		Active: trainingSheets.Active,
	}, nil
}
