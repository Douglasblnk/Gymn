package authService

import (
	userRepository "gymn/internal/repository/user"
	"gymn/internal/utils"
)

func ValidateEmail(email string) (bool, *utils.Error) {
	_, err := userRepository.FindUserByEmail(email)

	if err != nil {
		return false, err
	}

	return true, nil
}
