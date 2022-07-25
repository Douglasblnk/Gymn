package authRepository

import (
	"gymn/internal/database"
	"gymn/internal/exceptions"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func FindSessionByRefreshToken(refreshToken string) (*models.Session, *utils.Error) {
	session := &models.Session{}

	if err := database.DB.Where("refresh_token = ?", refreshToken).First(session).Error; err != nil {
		return nil, utils.Throw(exceptions.ErrSessionNotFound, 404)
	}

	return session, nil
}
