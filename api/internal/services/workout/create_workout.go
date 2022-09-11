package trainingSheetService

import (
	"gymn/internal/dto"
	"gymn/internal/models"
	workoutRepository "gymn/internal/repository/workout"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func CreateWorkout(userID int, data *schemas.RegisterWorkout) (*dto.WorkoutDTO, *utils.Error) {
	user, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return nil, err
	}

	workout := &models.Workout{
		Exercise:    data.Exercise,
		Description: data.Description,
		Repetitions: data.Repetitions,
		Series:      data.Series,
		RestTime:    data.RestTime,
		Weight:      data.Weight,
		Cadence:     data.Cadence,
	}

	if err := workoutRepository.CreateWorkout(user, workout); err != nil {
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
