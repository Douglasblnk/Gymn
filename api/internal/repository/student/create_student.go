package studentRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func CreateStudent(student *models.Student) *utils.Error {
	if err := database.DB.Select("*").Create(student).Error; err != nil {
		return utils.Throw(err.Error(), 404)
	}

	return nil
}
