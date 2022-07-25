package authService

import (
	"gymn/internal/dto"
	authRepository "gymn/internal/repository/auth"
	userRepository "gymn/internal/repository/user"
	"gymn/internal/security"
	"gymn/internal/utils"
)

func CreateAccessToken(refreshToken string) (*dto.AccessToken, *utils.Error) {
	session, err := authRepository.FindSessionByRefreshToken(refreshToken)

	if err != nil {
		return nil, err
	}

	user, err := userRepository.FindUserByID(session.UserID)

	if err != nil {
		return nil, err
	}

	jwtPayload := security.JWTPayload{}

	jwtPayload["user_uid"] = user.UID
	jwtPayload["session_uid"] = session.UID

	tokenData, err := security.EncryptJWT(jwtPayload)

	if err != nil {
		return nil, err
	}

	accessToken := &dto.AccessToken{
		Token:      tokenData.Token,
		Expiration: tokenData.Expiration,
	}

	return accessToken, nil
}
