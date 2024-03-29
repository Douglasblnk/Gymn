package studentService

import (
	studentRepository "gymn/internal/repository/student"
	"gymn/internal/utils"
)

func GetStudentCode(userID int, uid string) (map[string]string, *utils.Error) {
	student, err := studentRepository.GetStudentByUID(userID, uid)

	if err != nil {
		return nil, err
	}

	return map[string]string{"code": student.Code}, nil
}
