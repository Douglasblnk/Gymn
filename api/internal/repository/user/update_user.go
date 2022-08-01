package userRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func UpdateUser(user *models.User) *utils.Error {
	if err := database.DB.Save(user).Error; err != nil {
		return utils.Throw(err.Error(), 400)
	}

	return nil
}
