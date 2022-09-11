package trainingSheetRepository

import (
	"gymn/internal/database"
	"gymn/internal/utils"
)

func AssociateStudentTrainingSheet(studentID int, trainingSheetID int) *utils.Error {
	err := database.DB.Exec(
		`INSERT INTO public.student_training_sheets (student_id, training_sheet_id) VALUES ($1,$2);`, studentID, trainingSheetID,
	).Error

	if err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
