package trainingSheetRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func DeleteTrainingSheet(trainingSheet *models.TrainingSheet) *utils.Error {
	if err := database.DB.Delete(trainingSheet).Error; err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
