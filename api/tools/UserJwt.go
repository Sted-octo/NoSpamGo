package tools

import jwt "github.com/golang-jwt/jwt/v5"

//UserJWT is to parse info from jw tocken
type UserJwt struct {
	Email string `json:"ema"`
	jwt.RegisteredClaims
}
