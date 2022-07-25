package authService

import (
	"gymn/internal/models"
	authRepository "gymn/internal/repository/auth"
	"gymn/internal/utils"
	"os"
	"strconv"
)

func CreateSession(userId int) (*models.Session, *utils.Error) {
	refreshTokenLength, _ := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_LENGTH"), 10, 64)

	refreshToken := utils.GenerateRandomText(refreshTokenLength)

	session := &models.Session{
		UserID:       userId,
		RefreshToken: refreshToken,
	}

	if err := authRepository.CreateSession(session); err != nil {
		return nil, err
	}

	return session, nil
}
