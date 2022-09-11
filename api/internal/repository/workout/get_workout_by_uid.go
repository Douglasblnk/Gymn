package workoutRepository

import (
	"gymn/internal/database"
	"gymn/internal/exceptions"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetWorkoutByUID(userID int, uid string) (*models.Workout, *utils.Error) {
	workout := &models.Workout{}

	if err := database.DB.Where("uid = ? AND user_id = ?", uid, userID).First(workout).Error; err != nil {
		return nil, utils.Throw(exceptions.ErrWorkoutNotFound, 404)
	}

	return workout, nil
}
