package studentRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func DeleteStudent(student *models.Student) *utils.Error {
	if err := database.DB.Delete(student).Error; err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
