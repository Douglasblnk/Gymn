package workoutService

import (
	workoutRepository "gymn/internal/repository/workout"
	"gymn/internal/utils"
)

func DeleteWorkout(userID int, workoutID string) (string, *utils.Error) {
	workout, err := workoutRepository.GetWorkoutByUID(userID, workoutID)

	if err != nil {
		return "", err
	}

	if err := workoutRepository.DeleteWorkout(workout); err != nil {
		return "", err
	}

	return "deleted", nil
}
