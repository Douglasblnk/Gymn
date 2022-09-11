package workoutRepository

import (
	"gymn/internal/database"
	"gymn/internal/utils"
)

func AssociateWorkoutTrainingSheet(trainingSheetID int, workoutID int) *utils.Error {
	err := database.DB.Exec(
		`INSERT INTO public.training_sheets_workouts (training_sheet_id, workout_id) VALUES ($1,$2);`, trainingSheetID, workoutID,
	).Error

	if err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
