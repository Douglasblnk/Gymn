package workoutService

import (
	"gymn/internal/dto"
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	workoutRepository "gymn/internal/repository/workout"
	"gymn/internal/utils"
)

func GetTrainingSheetWorkouts(userID int, trainingSheetID string) ([]*dto.WorkoutDTO, *utils.Error) {
	trainingSheet, err := trainingSheetRepository.GetTrainingSheetByUID(userID, trainingSheetID)

	if err != nil {
		return nil, err
	}

	workoutsModel, err := workoutRepository.GetAllTrainingSheetsWorkouts(userID, trainingSheet.ID)

	if err != nil {
		return nil, err
	}

	var workouts []*dto.WorkoutDTO

	for _, workout := range workoutsModel {
		workoutsDTO := &dto.WorkoutDTO{
			Exercise:    workout.Exercise,
			Description: workout.Description,
			Repetitions: workout.Repetitions,
			Series:      workout.Series,
			RestTime:    workout.RestTime,
			Weight:      workout.Weight,
			Cadence:     workout.Cadence,
		}

		workouts = append(workouts, workoutsDTO)
	}

	return workouts, nil
}
