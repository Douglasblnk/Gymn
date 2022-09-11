package workoutRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func DeleteWorkout(workout *models.Workout) *utils.Error {
	if err := database.DB.Delete(workout).Error; err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
