package studentService

import (
	"gymn/internal/dto"
	studentRepository "gymn/internal/repository/student"
	"gymn/internal/utils"
)

func GetStudents() ([]*dto.StudentDTO, *utils.Error) {
	studentsModel, err := studentRepository.GetAllStudents()

	if err != nil {
		return nil, err
	}

	var students []*dto.StudentDTO

	for _, student := range studentsModel {
		studentDTO := &dto.StudentDTO{
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Birth:     student.Birth,
			Color:     student.Color,
			StartDate: student.StartDate,
			Weight:    student.Weight,
			Height:    student.Height,
		}

		students = append(students, studentDTO)
	}

	return students, nil
}
