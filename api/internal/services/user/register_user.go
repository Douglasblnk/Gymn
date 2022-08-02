package userService

import (
	"gymn/internal/dto"
	"gymn/internal/exceptions"
	"gymn/internal/models"
	userRepository "gymn/internal/repository/user"
	"gymn/internal/security"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func RegisterUser(data *schemas.RegisterUser) (*dto.UserDTO, *utils.Error) {
	userWithSameEmail, err := userRepository.FindUserByEmail(data.Email)

	if err != nil && err.Error != exceptions.ErrUserNotFound {
		return nil, err
	}

	if userWithSameEmail != nil {
		return nil, utils.Throw(exceptions.ErrUserAlreadyExists, 400)
	}

	hashedPassword, err := security.HashPassword(data.Password)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:        data.Name,
		Email:       data.Email,
		Password:    hashedPassword,
		Is_personal: data.IsPersonal,
	}

	err = userRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		Name:        user.Email,
		Email:       user.Email,
		Is_personal: user.Is_personal,
	}, nil
}
