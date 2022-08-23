package trainingSheetService

import (
	"gymn/internal/models"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func CreateTrainingSheet(userID int, data *schemas.RegisterTrainingSheet) (*models.TrainingSheet, *utils.Error) {
	user, err := trainingSheetRepository.FindUserByID(userID)

	if err != nil {
		return nil, err
	}

	if *user.Is_personal == false {
		return nil, utils.Throw(exceptions.ErrStudentNotPersonal, 403)
	}

	student := &models.Student{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Birth:     data.Birth,
		Color:     data.Color,
		Code:      utils.GenerateRandomText(24),
		StartDate: (func() string {
			if data.StartDate != nil {
				return *data.StartDate
			}

			return time.Now().Format("01-02-2006")
		})(),
		Weight: data.Weight,
		Height: data.Height,
	}

	if err := trainingSheetRepository.CreateStudent(user, student); err != nil {
		return nil, err
	}

	return &dto.StudentDTO{
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Birth:     student.Birth,
		Color:     student.Color,
		StartDate: student.StartDate,
		Weight:    student.Weight,
		Height:    student.Height,
	}, nil
}
