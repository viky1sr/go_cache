package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/viky1sr/go_cache.git/app/models"
)

type UserValidator struct {
	validate *validator.Validate
}

func NewUserValidator() *UserValidator {
	validate := validator.New()
	return &UserValidator{validate: validate}
}

func (v *UserValidator) Validate(user *models.User) error {
	if err := v.validate.Struct(user); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			switch e.Field() {
			case "Name":
				switch e.Tag() {
				case "required":
					return fmt.Errorf("Name is required")
				case "max":
					return fmt.Errorf("Name must be maximum %v characters", e.Param())
				}
			case "Email":
				switch e.Tag() {
				case "required":
					return fmt.Errorf("Email is required")
				case "email":
					return fmt.Errorf("Invalid email format")
				}
			case "Password":
				switch e.Tag() {
				case "required":
					return fmt.Errorf("Password is required")
				case "min":
					return fmt.Errorf("Password must be at least %v characters", e.Param())
				}
			}
		}
	}
	return nil
}
