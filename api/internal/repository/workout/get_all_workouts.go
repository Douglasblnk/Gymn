package workoutRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetAllWorkouts(userID int) ([]*models.Workout, *utils.Error) {
	var workout []*models.Workout

	query := database.DB.Order("created_at desc")

	if err := query.Where("user_id = ?", userID).Find(&workout).Error; err != nil {
		return nil, utils.Throw(err.Error(), 400)
	}

	return workout, nil
}
