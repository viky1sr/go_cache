package repositories

import "github.com/viky1sr/go_cache.git/app/models"

type AuthRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateToken(userID uint64) (*models.Token, error)
}
