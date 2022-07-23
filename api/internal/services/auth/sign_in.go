package authService

import (
	"fmt"
	"gymn/internal/dto"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func SignIn(data *schemas.Login) (*dto.SignInResponseDTO, *utils.Error) {
	user, err := ValidateUser(data.Email, data.Password)

	if err != nil {
		return nil, err
	}

	fmt.Println("user", user)

	session, err := CreateSession(user.ID)

	fmt.Println("session", session)
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
