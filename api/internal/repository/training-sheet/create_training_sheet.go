package studentRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func CreateTrainingSheet(user *models.User, student *models.TrainingSheet) *utils.Error {
	if err := database.DB.Model(user).Association("Students").Append(student); err != nil {
		return utils.Throw(err.Error(), 404)
	}

	return nil
}
