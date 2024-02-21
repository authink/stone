package inkstone

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	jwt.RegisteredClaims
	AppId     int `json:"appId"`
	AccountId int `json:"accountId"`
}
