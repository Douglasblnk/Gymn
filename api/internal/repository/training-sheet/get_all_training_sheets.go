package trainingSheetRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetAllTrainingSheets(userID int) ([]*models.TrainingSheet, *utils.Error) {
	var trainingSheet []*models.TrainingSheet

	query := database.DB.Order("created_at desc")

	if err := query.Where("user_id = ?", userID).Find(&trainingSheet).Error; err != nil {
		return nil, utils.Throw(err.Error(), 400)
	}

	return trainingSheet, nil
}
