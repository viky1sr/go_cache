package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/viky1sr/go_cache.git/config"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the hashed password using bcrypt
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPasswordHash checks if the provided password matches the hashed password
func CheckPasswordHash(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

// GenerateToken generates a JWT token for the provided user ID
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.JWTSecret)) // convert to []byte
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// InvalidateToken invalidates the provided JWT token
func InvalidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return jwt.ErrInvalidKey
	}

	claims["exp"] = time.Now().Unix()
	token.Claims = claims

	return nil
}
