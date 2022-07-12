package authService

import (
	"gymn/internal/dto"
)

func ValidateUser(email string, password string) (*dto.UserDTO, error) {
	// user, err := userRepository.FindUserByUsername(email)

	// if err != nil {
	// 	return nil, err
	// }

	// passwordMatches := securityUtils.CompareHash(user.Password, password)

	// if !passwordMatches {
	// 	return nil, exceptions.ErrAuthorizationFailed
	// }

	// userDTO := &dto.UserDTO{
	// 	Name:     user.Name,
	// 	Username: user.Username,
	// }

	return nil, nil
}
