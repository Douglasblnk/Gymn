package workoutRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetAllTrainingSheetsWorkouts(userID int, trainingSheetID int) ([]*models.Workout, *utils.Error) {
	trainingSheet := &models.TrainingSheet{
		ID: trainingSheetID,
	}

	var workouts []*models.Workout

	database.DB.Model(trainingSheet).Association("Workout").Find(&workouts)

	return workouts, nil
}
