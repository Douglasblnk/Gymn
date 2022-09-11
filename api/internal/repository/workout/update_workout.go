package workoutRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func UpdateWorkout(workout *models.Workout) *utils.Error {
	if err := database.DB.Save(workout).Error; err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
