package studentService

import (
	"gymn/internal/dto"
	"gymn/internal/models"
	studentRepository "gymn/internal/repository/student"
	"gymn/internal/utils"
	"gymn/v1/schemas"
	"time"
)

func CreateStudent(data *schemas.RegisterStudent) (*dto.StudentDTO, *utils.Error) {
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

	if err := studentRepository.CreateStudent(student); err != nil {
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
