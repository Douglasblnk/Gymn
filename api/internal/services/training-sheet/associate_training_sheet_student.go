package trainingSheetService

import (
	"gymn/internal/exceptions"
	studentRepository "gymn/internal/repository/student"
	trainingSheetRepository "gymn/internal/repository/training-sheet"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func AssociateTrainingSheetStudent(userID int, data *schemas.StudentTrainingSheet) (bool, *utils.Error) {
	_, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return false, err
	}

	trainingSheet, err := trainingSheetRepository.GetTrainingSheetByUID(userID, data.TrainingSheetID)

	if err != nil {
		return false, err
	}

	if !trainingSheet.Active {
		return false, utils.Throw(exceptions.ErrTrainingSheetInactive, 403)
	}

	student, err := studentRepository.GetStudentByUID(userID, data.StudentID)

	if err != nil {
		return false, err
	}

	if err = trainingSheetRepository.AssociateTrainingSheetStudent(student.ID, trainingSheet.ID); err != nil {
		return false, err
	}

	return true, nil
}
