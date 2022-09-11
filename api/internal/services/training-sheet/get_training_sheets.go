package trainingSheetService

import (
	"gymn/internal/dto"
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	"gymn/internal/utils"
)

func GetTrainingSheets(userID int) ([]*dto.TrainingSheetDTO, *utils.Error) {
	trainingSheetsModel, err := trainingSheetRepository.GetAllTrainingSheets(userID)

	if err != nil {
		return nil, err
	}

	var trainingSheets []*dto.TrainingSheetDTO

	for _, trainingSheet := range trainingSheetsModel {
		trainingSheetsDTO := &dto.TrainingSheetDTO{
			Name:   trainingSheet.Name,
			Active: trainingSheet.Active,
		}

		trainingSheets = append(trainingSheets, trainingSheetsDTO)
	}

	return trainingSheets, nil
}
