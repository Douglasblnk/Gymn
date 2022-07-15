package userRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func CreateUser(user *models.User) *utils.Error {
	if err := database.DB.Select("*").Create(user).Error; err != nil {
		return utils.Throw(err.Error(), 404)
	}

	return nil
}
