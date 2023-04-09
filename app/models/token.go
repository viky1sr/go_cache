package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func (c *Claims) GenerateToken(secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (c *Claims) GetExpirationTime() time.Time {
	return time.Unix(c.ExpiresAt, 0)
}
