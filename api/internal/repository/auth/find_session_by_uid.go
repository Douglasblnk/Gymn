package authRepository

import (
	"gymn/internal/database"
	"gymn/internal/exceptions"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func FindSessionByUID(uid string) (*models.Session, *utils.Error) {
	session := &models.Session{}

	if err := database.DB.Where("uid = ?", uid).First(session).Error; err != nil {
		return nil, utils.Throw(exceptions.ErrSessionNotFound, 404)
	}

	return session, nil
}
