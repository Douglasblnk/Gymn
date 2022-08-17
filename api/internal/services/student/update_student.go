package studentService

import (
	"gymn/internal/dto"
	"gymn/internal/exceptions"
	studentRepository "gymn/internal/repository/student"
	userRepository "gymn/internal/repository/user"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func UpdateStudent(userID int, id string, data *schemas.UpdateStudent) (*dto.StudentDTO, *utils.Error) {
	user, err := userRepository.FindUserByID(userID)

	if err != nil {
		return nil, err
	}

	if *user.Is_personal == false {
		return nil, utils.Throw(exceptions.ErrStudentNotPersonal, 403)
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
