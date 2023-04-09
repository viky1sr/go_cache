package services

import (
	"errors"
	"fmt"
	"github.com/viky1sr/go_cache.git/app/models"
	"github.com/viky1sr/go_cache.git/app/repositories"
	"github.com/viky1sr/go_cache.git/app/validators"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id int, user *models.User) error
	DeleteUser(id int) error
}

type userService struct {
	userRepo      repositories.UserRepository
	userValidator *validators.UserValidator
}

func NewUserService(userRepo repositories.UserRepository, userValidator *validators.UserValidator) UserService {
	return &userService{userRepo: userRepo, userValidator: userValidator}
}

func (s *userService) GetAllUsers() ([]*models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	return s.userRepo.GetUserByID(id)
}

func (s *userService) CreateUser(user *models.User) error {
	err := s.userValidator.Validate(user)
	if err != nil {
		return err
	}

	// Check if email already exists in the database
	existingUserByEmail, err := s.userRepo.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUserByEmail != nil {
		return fmt.Errorf("Email already exists")
	}

	user.Password = hashPassword(user.Password)

	// Create the user
	err = s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUser(id int, user *models.User) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}

	err := s.userValidator.Validate(user)
	if err != nil {
		return err
	}

	// Check user in the database
	_, err = s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	// Check if email already exists in the database, except for the current user
	existingUserByEmail, err := s.userRepo.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUserByEmail != nil && existingUserByEmail.ID != uint(id) {
		return fmt.Errorf("Email already exists")
	}

	user.Password = hashPassword(user.Password)

	return s.userRepo.UpdateUser(id, user)
}

func (s *userService) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}

	return s.userRepo.DeleteUser(id)
}

func hashPassword(password string) string {
	// Generate the hashed password using the bcrypt algorithm
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	// Return the hashed password as a string
	return string(hashedPassword)
}
