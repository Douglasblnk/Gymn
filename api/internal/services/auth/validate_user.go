package authService

import (
	"gymn/internal/exceptions"
	"gymn/internal/models"
	userRepository "gymn/internal/repository/user"
	"gymn/internal/security"
	"gymn/internal/utils"
)

func ValidateUser(email string, password string) (*models.User, *utils.Error) {
	user, err := userRepository.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	passwordMatches := security.CompareHash(user.Password, password)

	if !passwordMatches {
		return nil, utils.Throw(exceptions.ErrAuthorizationFailed, 401)
	}

	return user, nil
}
