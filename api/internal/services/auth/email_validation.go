package authService

import (
	userRepository "gymn/internal/repository/user"
	"gymn/internal/utils"
	"strings"
)

func ValidateEmail(email string) (bool, *utils.Error) {
	_, err := userRepository.FindUserByEmail(strings.ToLower(email))

	if err != nil {
		return false, err
	}

	return true, nil
}
