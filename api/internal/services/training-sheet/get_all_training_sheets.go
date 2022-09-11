package trainingSheetService

import (
	"gymn/internal/dto"
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	"gymn/internal/utils"
)

func GetAllTrainingSheets(userID int) ([]*dto.TrainingSheetDTO, *utils.Error) {
	trainingSheetsModel, err := trainingSheetRepository.GetAllTrainingSheets(userID)

	if err != nil {
		return nil, err
	}

	var trainingSheets []*dto.TrainingSheetDTO

	for _, trainingSheet := range trainingSheetsModel {
		trainingSheetDTO := &dto.TrainingSheetDTO{
			Name:   trainingSheet.Name,
			Active: trainingSheet.Active,
		}

		trainingSheets = append(trainingSheets, trainingSheetDTO)
	}

	return trainingSheets, nil
}
