package authService

import (
	"gymn/internal/dto"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func SignIn(data *schemas.Login) (*dto.SignInResponseDTO, *utils.Error) {
	user, err := ValidateUser(data.Email, data.Password)

	if err != nil {
		return nil, err
	}

	session, err := CreateSession(user)

	if err != nil {
		return nil, err
	}

	accessToken, err := CreateAccessToken(session.RefreshToken)

	if err != nil {
		return nil, err
	}

	response := &dto.SignInResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: session.RefreshToken,
		User: &dto.UserDTO{
			Name:        user.Name,
			Email:       user.Email,
			Is_personal: user.Is_personal,
		},
	}

	return response, nil
}
