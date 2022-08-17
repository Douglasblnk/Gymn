package authService

import (
	"gymn/internal/exceptions"
	"gymn/internal/models"
	authRepository "gymn/internal/repository/auth"
	"gymn/internal/security"
	"gymn/internal/utils"
)

func ValidateAccessToken(accessToken string) (*models.Session, *utils.Error) {
	decrypted, decryptErr := security.DecryptJWT(accessToken)

	if decryptErr != nil {
		return nil, utils.Throw(exceptions.ErrAuthorizationFailed, 401)
	}

	sessionUID := decrypted["session_uid"].(string)

	session, err := authRepository.FindSessionByUID(sessionUID)

	if err != nil {
		return nil, utils.Throw(exceptions.ErrSessionNotFound, 404)
	}

	return session, nil
}
