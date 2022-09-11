package trainingSheetRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetAllStudentTrainingSheets(userID int, studentID int) ([]*models.TrainingSheet, *utils.Error) {
	student := &models.Student{
		ID: studentID,
	}

	var trainingSheet []*models.TrainingSheet

	database.DB.Model(student).Association("TrainingSheet").Find(&trainingSheet)

	return trainingSheet, nil
}
