package authService

import (
	"gymn/internal/exceptions"
	authRepository "gymn/internal/repository/auth"
	"gymn/internal/security"
	"gymn/internal/utils"
)

func ValidateAccessToken(accessToken string) *utils.Error {
	decrypted, decryptErr := security.DecryptJWT(accessToken)

	if decryptErr != nil {
		return utils.Throw(decryptErr.Error(), 401)
	}

	sessionUID := decrypted["session_uid"].(string)

	_, err := authRepository.FindSessionByUID(sessionUID)

	if err != nil {
		return utils.Throw(exceptions.ErrSessionNotFound, 404)
	}

	return nil
}
