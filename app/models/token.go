package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}
