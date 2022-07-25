package security

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func DecryptJWT(token string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}
