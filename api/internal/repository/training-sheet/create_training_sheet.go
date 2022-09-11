package trainingSheetRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func CreateTrainingSheet(user *models.User, trainingSheet *models.TrainingSheet) *utils.Error {
	if err := database.DB.Model(user).Association("TrainingSheet").Append(trainingSheet); err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
