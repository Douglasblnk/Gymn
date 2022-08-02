package studentService

import (
	"gymn/internal/dto"
	studentRepository "gymn/internal/repository/student"
	"gymn/internal/utils"
)

func GetStudent(id string) (*dto.StudentDTO, *utils.Error) {
	student, err := studentRepository.GetStudentByUID(id)

	if err != nil {
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
