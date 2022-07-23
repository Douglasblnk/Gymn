package authService

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
	"os"
	"strconv"
)

func CreateSession(userId int) (*models.Session, *utils.Error) {
	refreshTokenLength, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_LENGTH"), 10, 64)

	if err != nil {
		return nil, utils.Throw(err.Error(), 500)
	}

	refreshToken := utils.GenerateRandomText(refreshTokenLength)

	session := &models.Session{
		UserID:       userId,
		RefreshToken: refreshToken,
	}

	if err := database.DB.Select("*").Create(session).Error; err != nil {
		return nil, utils.Throw(err.Error(), 500)
	}

	return session, nil
}
