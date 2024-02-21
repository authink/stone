package inkstone

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

func GenerateToken(issuer, key string, duration time.Duration, appId uint32, appName string, accountId uint32, email, uuid string) (string, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer: issuer,
				Audience: jwt.ClaimStrings{
					appName,
				},
				Subject:   email,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration * time.Second)),
				ID:        uuid,
			},
			AppId:     int(appId),
			AccountId: int(accountId),
		},
	)

	return t.SignedString([]byte(key))
}

func VerifyToken(key string, accessToken string) (jwtClaims *JwtClaims, err error) {
	jwtClaims = &JwtClaims{}

	_, err = jwt.ParseWithClaims(
		accessToken,
		jwtClaims,
		func(token *jwt.Token) (any, error) {
			return []byte(key), nil
		},
	)
	return
}

func GenerateUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
