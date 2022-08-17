package studentService

import (
	studentRepository "gymn/internal/repository/student"
	"gymn/internal/utils"
)

func DeleteStudent(userID int, studentID string) (string, *utils.Error) {
	student, err := studentRepository.GetStudentByUID(userID, studentID)

	if err != nil {
		return "", err
	}

	if err := studentRepository.DeleteStudent(student); err != nil {
		return "", err
	}

	return "deleted", nil
}
