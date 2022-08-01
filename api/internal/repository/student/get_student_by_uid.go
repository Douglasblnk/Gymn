package studentRepository

import (
	"gymn/internal/database"
	"gymn/internal/exceptions"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetStudentByUID(uid string) (*models.Student, *utils.Error) {
	student := &models.Student{}

	if err := database.DB.Where("uid = ?", uid).First(student).Error; err != nil {
		return nil, utils.Throw(exceptions.ErrStudentNotFound, 404)
	}

	return student, nil
}
