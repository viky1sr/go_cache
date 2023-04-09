package services

import (
	"errors"
	"github.com/viky1sr/go_cache.git/app/models"
	"github.com/viky1sr/go_cache.git/app/repositories"
	"github.com/viky1sr/go_cache.git/app/utils"
)

type AuthService interface {
	Login(email, password string) (*models.Token, error)
	Logout(tokenString string) error
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Login(email, password string) (*models.Token, error) {
	// Get user by email
	user, _ := s.userRepo.FindByEmail(email)
	if user == nil {
		return nil, errors.New("Data not found")
	}

	// Check password
	if !utils.CheckPasswordHash([]byte(password), []byte(user.Password)) {
		return nil, errors.New("invalid login credentials")
	}

	//Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.Token{AccessToken: token}, nil
}

func (s *authService) Logout(tokenString string) error {
	// Invalidate token
	err := utils.InvalidateToken(tokenString)
	if err != nil {
		return err
	}

	return nil
}
