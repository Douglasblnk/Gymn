package trainingSheetRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func UpdateTrainingSheet(trainingSheet *models.TrainingSheet) *utils.Error {
	if err := database.DB.Save(trainingSheet).Error; err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
