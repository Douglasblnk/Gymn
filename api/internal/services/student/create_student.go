package userService

import (
	"gymn/internal/models"
	studentRepository "gymn/internal/repository/student"
	"gymn/internal/utils"
	"gymn/v1/schemas"
	"time"
)

func CreateStudent(data *schemas.RegisterStudent) (*models.Student, *utils.Error) {
	student := &models.Student{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Birth:     data.Birth,
		Color:     data.Color,
		StartDate: time.Now().Format("01-02-2006"),
		Weight:    data.Weight,
		Height:    data.Height,
	}

	if err := studentRepository.CreateStudent(student); err != nil {
		return nil, err
	}

	return student, nil
}
