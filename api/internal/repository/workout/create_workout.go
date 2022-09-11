package trainingSheetRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func CreateWorkout(user *models.User, workout *models.Workout) *utils.Error {
	if err := database.DB.Model(user).Association("Workout").Append(workout); err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
