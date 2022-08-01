package studentService

import (
	"gymn/internal/models"
	studentRepository "gymn/internal/repository/student"
	"gymn/internal/utils"
)

func GetStudents() ([]*models.Student, *utils.Error) {
	students, err := studentRepository.GetStudents()

	if err != nil {
		return nil, err
	}

	return students, nil
}
