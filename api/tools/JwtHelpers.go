package tools

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func IsTokenValide(tokenString string, cryptoKey string) (*UserJwt, error) {
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {

		return nil, errors.New("invalid authorization format")
	}

	splitedTokenString := parts[1]

	token, err := jwt.ParseWithClaims(splitedTokenString, &UserJwt{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(cryptoKey), nil
	})

	if claims, ok := token.Claims.(*UserJwt); ok && token.Valid {
		expire, err := claims.GetExpirationTime()
		if err != nil {
			return nil, err
		}
		if expire.Compare(time.Now()) > 0 {
			return claims, nil
		}
	}

	return nil, err
}

func RenewToken(userJwt *UserJwt, cryptoKey string) string {
	day, _ := time.ParseDuration("48h00m")
	userJwt.ExpiresAt = jwt.NewNumericDate(time.Now().Add(day))
	userJwt.IssuedAt = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userJwt)

	tokenString, error := token.SignedString([]byte(cryptoKey))
	if error != nil {
		fmt.Println(error)
		return ""
	}
	return tokenString
}

func GenerateToken(mail string, secretKey string) (string, error) {
	day, _ := time.ParseDuration("48h00m")
	claims := &UserJwt{
		Email: mail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(day)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
