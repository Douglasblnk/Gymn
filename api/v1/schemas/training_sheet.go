package schemas

type RegisterTrainingSheet struct {
	Name string `json:"name" validate:"required"`
}

type UpdateTrainingSheet struct {
	Name   *string `json:"name"`
	Active *bool   `json:"active"`
}

type StudentTrainingSheet struct {
	StudentID       string `json:"studentId"`
	TrainingSheetID string `json:"trainingSheetId"`
}
