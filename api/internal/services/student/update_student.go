package studentService

import (
	"gymn/internal/dto"
	studentRepository "gymn/internal/repository/student"
	userService "gymn/internal/services/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func UpdateStudent(userID int, id string, data *schemas.UpdateStudent) (*dto.StudentDTO, *utils.Error) {
	_, err := userService.ValidateUserPersonal(userID)

	if err != nil {
		return nil, err
	}

	student, err := studentRepository.GetStudentByUID(userID, id)

	if err != nil {
		return nil, err
	}

	if data.FirstName != nil {
		student.FirstName = *data.FirstName
	}

	if data.LastName != nil {
		student.LastName = *data.LastName
	}

	if data.Birth != nil {
		student.Birth = data.Birth
	}

	if data.Color != nil {
		student.Color = data.Color
	}

	if data.StartDate != nil {
		student.StartDate = *data.StartDate
	}

	if data.Weight != nil {
		student.Weight = *data.Weight
	}

	if data.Height != nil {
		student.Height = *data.Height
	}

	if err = studentRepository.UpdateStudent(student); err != nil {
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
