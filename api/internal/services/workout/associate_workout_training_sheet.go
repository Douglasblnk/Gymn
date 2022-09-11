package workoutService

import (
	"gymn/internal/exceptions"
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	workoutRepository "gymn/internal/repository/workout"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func AssociateWorkoutTrainingSheet(userID int, data *schemas.WorkoutTrainingSheet) (bool, *utils.Error) {
	_, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return false, err
	}

	workout, err := workoutRepository.GetWorkoutByUID(userID, data.WorkoutID)

	if err != nil {
		return false, err
	}

	trainingSheet, err := trainingSheetRepository.GetTrainingSheetByUID(userID, data.TrainingSheetID)

	if err != nil {
		return false, err
	}

	if !trainingSheet.Active {
		return false, utils.Throw(exceptions.ErrTrainingSheetInactive, 403)
	}

	if err = workoutRepository.AssociateWorkoutTrainingSheet(trainingSheet.ID, workout.ID); err != nil {
		return false, err
	}

	return true, nil
}
