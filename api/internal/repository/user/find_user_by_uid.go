package userRepository

import (
	"gymn/internal/database"
	"gymn/internal/exceptions"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func FindUserByUID(uid string) (*models.User, *utils.Error) {
	user := &models.User{}

	if err := database.DB.Where("uid = ?", uid).First(user).Error; err != nil {
		return nil, utils.Throw(exceptions.ErrUserNotFound, 404)
	}

	return user, nil
}
