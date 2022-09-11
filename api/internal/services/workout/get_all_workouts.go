package workoutService

import (
	"gymn/internal/dto"
	workoutRepository "gymn/internal/repository/workout"
	"gymn/internal/utils"
)

func GetAllWorkouts(userID int) ([]*dto.WorkoutDTO, *utils.Error) {
	workoutModel, err := workoutRepository.GetAllWorkouts(userID)

	if err != nil {
		return nil, err
	}

	var workouts []*dto.WorkoutDTO

	for _, workout := range workoutModel {
		workoutDTO := &dto.WorkoutDTO{
			Exercise:    workout.Exercise,
			Description: workout.Description,
			Repetitions: workout.Repetitions,
			Series:      workout.Series,
			RestTime:    workout.RestTime,
			Weight:      workout.Weight,
			Cadence:     workout.Cadence,
		}

		workouts = append(workouts, workoutDTO)
	}

	return workouts, nil
}
