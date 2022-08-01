package studentRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func UpdateStudent(student *models.Student) *utils.Error {
	if err := database.DB.Save(student).Error; err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
