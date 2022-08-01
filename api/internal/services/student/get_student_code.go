package studentService

import (
	studentRepository "gymn/internal/repository/student"
	"gymn/internal/utils"
)

func GetStudentCode(uid string) (map[string]string, *utils.Error) {
	student, err := studentRepository.GetStudentByUID(uid)

	if err != nil {
		return nil, err
	}

	return map[string]string{"code": student.Code}, nil
}
