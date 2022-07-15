package userRepository

import (
	"gymn/internal/database"
	"gymn/internal/exceptions"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func FindUserByEmail(email string) (*models.User, *utils.Error) {
	user := &models.User{}

	if err := database.DB.Where("email = ?", email).First(user).Error; err != nil {
		return nil, utils.Throw(exceptions.ErrUserNotFound, 404)
	}

	return user, nil
}
