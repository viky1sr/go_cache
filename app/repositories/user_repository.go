package repositories

import (
	"github.com/viky1sr/go_cache.git/app/models"
)

type UserRepository interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id int, user *models.User) error
	DeleteUser(id int) error
}
