package authService

import (
	"gymn/internal/dto"
	authRepository "gymn/internal/repository/auth"
	"gymn/internal/utils"
)

func RefreshAccessToken(refreshToken string) (*dto.RefreshAccessTokenDTO, *utils.Error) {
	accessToken, err := CreateAccessToken(refreshToken)

	if err != nil {
		return nil, err
	}

	session, err := authRepository.FindSessionByRefreshToken(refreshToken)

	if err != nil {
		return nil, err
	}

	session.RefreshToken = utils.GenerateRandomText(
		utils.GetIntEnv("REFRESH_TOKEN_LENGTH"),
	)

	if err = authRepository.UpdateSession(session); err != nil {
		return nil, err
	}

	response := &dto.RefreshAccessTokenDTO{
		AccessToken:  accessToken,
		RefreshToken: session.RefreshToken,
	}

	return response, nil
}
