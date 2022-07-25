package security

import (
	"gymn/internal/dto"
	"gymn/internal/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTPayload map[string]interface{}

func EncryptJWT(payload JWTPayload) (*dto.EncryptJWTResponseDTO, *utils.Error) {
	claims := jwt.MapClaims{}

	expiration := time.Now().Add(time.Hour * time.Duration(utils.GetIntEnv("JWT_EXPIRATION_HOURS")))

	claims["exp"] = expiration.Unix()

	for key, value := range payload {
		claims[key] = value
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, utils.Throw(err.Error(), 500)
	}

	response := &dto.EncryptJWTResponseDTO{
		Token:      token,
		Expiration: expiration,
	}

	return response, nil
}
