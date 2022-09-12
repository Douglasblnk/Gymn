package authRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func CreateSession(user *models.User, session *models.Session) *utils.Error {
	if err := database.DB.Model(user).Association("Session").Append(session); err != nil {
		return utils.Throw(err.Error(), 500)
	}

	return nil
}
