package authRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func CreateSession(session *models.Session) *utils.Error {
	if err := database.DB.Select("*").Create(session).Error; err != nil {
		return utils.Throw(err.Error(), 500)
	}

	return nil
}
