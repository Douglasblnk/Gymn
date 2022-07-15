package security

import (
	"gymn/internal/utils"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *utils.Error) {
	bcost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	if err != nil {
		return "", utils.Throw(err.Error(), 500)
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcost,
	)

	if err != nil {
		return "", utils.Throw(err.Error(), 500)
	}

	return string(hash), nil
}
