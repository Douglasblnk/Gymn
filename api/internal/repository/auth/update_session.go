package authRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func UpdateSession(session *models.Session) *utils.Error {
	if err := database.DB.Save(session).Error; err != nil {
		return utils.Throw(err.Error(), 500)
	}

	return nil
}
