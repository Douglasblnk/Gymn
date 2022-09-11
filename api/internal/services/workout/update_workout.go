package workoutService

import (
	"gymn/internal/dto"
	workoutRepository "gymn/internal/repository/workout"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func UpdateWorkout(userID int, id string, data *schemas.UpdateWorkout) (*dto.WorkoutDTO, *utils.Error) {
	_, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return nil, err
	}

	workout, err := workoutRepository.GetWorkoutByUID(userID, id)

	if err != nil {
		return nil, err
	}

	if data.Exercise != nil {
		workout.Exercise = *data.Exercise
	}

	if data.Description != nil {
		workout.Description = *data.Description
	}

	if data.Repetitions != nil {
		workout.Repetitions = *data.Repetitions
	}

	if data.Series != nil {
		workout.Series = *data.Series
	}

	if data.RestTime != nil {
		workout.RestTime = *data.RestTime
	}

	if data.Weight != nil {
		workout.Weight = *data.Weight
	}

	if data.Cadence != nil {
		workout.Cadence = *data.Cadence
	}

	if err = workoutRepository.UpdateWorkout(workout); err != nil {
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
