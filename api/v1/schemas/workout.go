package schemas

type RegisterWorkout struct {
	Exercise    string  `json:"exercise" validate:"required"`
	Description string  `json:"description"`
	Repetitions string  `json:"repetitions" validate:"required"`
	Series      string  `json:"series" validate:"required"`
	RestTime    string  `json:"restTime"`
	Weight      float32 `json:"weight"`
	Cadence     string  `json:"cadence"`
}

type UpdateWorkout struct {
	Exercise    *string  `json:"exercise"`
	Description *string  `json:"description"`
	Repetitions *string  `json:"repetitions"`
	Series      *string  `json:"series"`
	RestTime    *string  `json:"restTime"`
	Weight      *float32 `json:"weight"`
	Cadence     *string  `json:"cadence"`
}

type WorkoutTrainingSheet struct {
	TrainingSheetID string `json:"trainingSheetId"`
	WorkoutID       string `json:"workoutId"`
}
