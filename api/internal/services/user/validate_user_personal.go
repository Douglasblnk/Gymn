package userService

import (
	"gymn/internal/exceptions"
	"gymn/internal/models"
	userRepository "gymn/internal/repository/user"
	"gymn/internal/utils"
)

func ValidateUserPersonal(userID int) (*models.User, *utils.Error) {
	user, err := userRepository.FindUserByID(userID)

	if err != nil {
		return nil, err
	}

	if *user.Is_personal == false {
		return nil, utils.Throw(exceptions.ErrUserNotPersonal, 403)
	}

	return user, nil
}
