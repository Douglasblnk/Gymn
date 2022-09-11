package trainingSheetService

import (
	"gymn/internal/dto"
	studentRepository "gymn/internal/repository/student"
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	"gymn/internal/utils"
)

func GetStudentTrainingSheets(userID int, studentID string) ([]*dto.TrainingSheetDTO, *utils.Error) {
	student, err := studentRepository.GetStudentByUID(userID, studentID)

	if err != nil {
		return nil, err
	}

	trainingSheetsModel, err := trainingSheetRepository.GetAllStudentTrainingSheets(userID, student.ID)

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
