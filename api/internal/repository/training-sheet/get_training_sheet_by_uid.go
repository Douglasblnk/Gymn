package trainingSheetRepository

import (
	"gymn/internal/database"
	"gymn/internal/exceptions"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetTrainingSheetByUID(userID int, uid string) (*models.TrainingSheet, *utils.Error) {
	trainingSheet := &models.TrainingSheet{}

	if err := database.DB.Where("uid = ? AND user_id = ?", uid, userID).First(trainingSheet).Error; err != nil {
		return nil, utils.Throw(exceptions.ErrTrainingSheetNotFound, 404)
	}

	return trainingSheet, nil
}
