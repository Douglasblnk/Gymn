package security

import (
	"gymn/internal/utils"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bcost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	if err != nil {
		return "", utils.Throw(err, 404)
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
