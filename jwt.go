package inkstone

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	AppId     int `json:"appId"`
	AccountId int `json:"accountId"`
}

func NewJwtClaims(id, issuer, appName, email string, duration time.Duration, appId, accountId uint32) *JwtClaims {
	return &JwtClaims{
		jwt.RegisteredClaims{
			Issuer: issuer,
			Audience: jwt.ClaimStrings{
				appName,
			},
			Subject:   email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration * time.Second)),
			ID:        id,
		},
		int(appId),
		int(accountId),
	}
}

func GenerateToken(key string, jwtClaims *JwtClaims) (string, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwtClaims,
	)

	return t.SignedString([]byte(key))
}

func VerifyToken(key string, accessToken string) (jwtClaims *JwtClaims, err error) {
	jwtClaims = new(JwtClaims)

	_, err = jwt.ParseWithClaims(
		accessToken,
		jwtClaims,
		func(token *jwt.Token) (any, error) {
			return []byte(key), nil
		},
	)
	return
}
