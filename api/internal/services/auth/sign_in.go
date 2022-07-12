package authService

import (
	"gymn/internal/dto"
	"gymn/v1/schemas"
)

func SignIn(data *schemas.Login) (*dto.SignInResponseDTO, error) {
	// user, err := ValidateUser(data.Email, data.Password)

	// if err != nil {
	// 	return nil, err
	// }

	// session, err := CreateSession(user.Username)

	// if err != nil {
	// 	return nil, err
	// }

	// accessToken, err := CreateAccessToken(session.RefreshToken)

	// if err != nil {
	// 	return nil, err
	// }

	// response := &dto.SignInResponseDTO{
	// 	AccessToken:  accessToken,
	// 	RefreshToken: session.RefreshToken,
	// 	User:         user,
	// }

	return nil, nil
}
