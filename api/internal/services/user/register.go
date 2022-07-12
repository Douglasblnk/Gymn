package userService

import (
	"gymn/internal/dto"
	"gymn/internal/exceptions"
	userRepository "gymn/internal/repository/user"
	"gymn/internal/security"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func RegisterUser(data *schemas.User) (*dto.UserDTO, *utils.Error) {
	userWithSameEmail, err := userRepository.FindUserByEmail(data.Email)

	if err != nil && err.Error != exceptions.ErrUserNotFound {
		return nil, err
	}

	if userWithSameEmail != nil {
		return nil, utils.Throw(exceptions.ErrUserAlreadyExists, 400)
	}

	hashedPassowrd, err := security.HashPassword(data.Password)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
