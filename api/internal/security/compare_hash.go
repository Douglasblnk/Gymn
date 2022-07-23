package security

import "golang.org/x/crypto/bcrypt"

func CompareHash(hash string, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))

	return err == nil
}
