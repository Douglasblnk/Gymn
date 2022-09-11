package workoutService

import (
	"gymn/internal/dto"
	workoutRepository "gymn/internal/repository/workout"
	"gymn/internal/utils"
)

func GetWorkout(userID int, workoutID string) (*dto.WorkoutDTO, *utils.Error) {
	workout, err := workoutRepository.GetWorkoutByUID(userID, workoutID)

	if err != nil {
		return nil, err
	}

	return &dto.WorkoutDTO{
		Exercise:    workout.Exercise,
		Description: workout.Description,
		Repetitions: workout.Repetitions,
		Series:      workout.Series,
		RestTime:    workout.RestTime,
		Weight:      workout.Weight,
		Cadence:     workout.Cadence,
	}, nil
}
