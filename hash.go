package inkstone

import (
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func Sha256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

func HashPassword(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return
	}

	hash = string(bytes)
	return
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}
