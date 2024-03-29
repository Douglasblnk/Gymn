package studentRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func CreateStudent(user *models.User, student *models.Student) *utils.Error {
	if err := database.DB.Model(user).Association("Students").Append(student); err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
