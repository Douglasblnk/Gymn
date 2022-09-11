package schemas

type RegisterWorkout struct {
	Exercise    string  `json:"exercise" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Repetitions string  `json:"repetitions" validate:"required"`
	Series      int     `json:"series" validate:"required"`
	RestTime    int     `json:"restTime" validate:"required"`
	Weight      float32 `json:"weight" validate:"required"`
	Cadence     string  `json:"cadence" validate:"required"`
}

type UpdateWorkout struct {
	Exercise    *string  `json:"exercise"`
	Description *string  `json:"description"`
	Repetitions *string  `json:"repetitions"`
	Series      *int     `json:"series"`
	RestTime    *int     `json:"restTime"`
	Weight      *float32 `json:"weight"`
	Cadence     *string  `json:"cadence"`
}

type TrainingSheetWorkout struct {
	TrainingSheetId string `json:"trainingSheetId"`
	WorkoutId       string `json:"workoutId"`
}
