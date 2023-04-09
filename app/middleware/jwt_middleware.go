package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/viky1sr/go_cache.git/app/models"
	"github.com/viky1sr/go_cache.git/app/traits"
	"github.com/viky1sr/go_cache.git/config"
)

// JWTMiddleware is middleware to check JWT token
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header value
		tokenHeader := r.Header.Get("Authorization")

		// Check if the Authorization header is present
		if tokenHeader == "" {
			responseTrait := traits.ResponseTrait{}
			responseTrait.RespondWithFailure(w, http.StatusUnauthorized, "Missing Authorization Header")
			return
		}

		// The Authorization header format should be "Bearer {token}"
		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			responseTrait := traits.ResponseTrait{}
			responseTrait.RespondWithFailure(w, http.StatusUnauthorized, "Invalid Authorization Header Format")
			return
		}

		// Get the token
		tokenString := tokenParts[1]

		// Validate the token
		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			responseTrait := traits.ResponseTrait{}
			responseTrait.RespondWithFailure(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Check if the token is valid
		claims, ok := token.Claims.(*models.Claims)
		if !ok || !token.Valid {
			responseTrait := traits.ResponseTrait{}
			responseTrait.RespondWithFailure(w, http.StatusUnauthorized, "Invalid Token")
			return
		}

		// Add the user ID to the request context
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
