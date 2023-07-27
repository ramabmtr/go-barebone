package entity

import "github.com/golang-jwt/jwt/v5"

const (
	CtxAuthInfo = "ctx::authInfo"
)

type JWTCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
