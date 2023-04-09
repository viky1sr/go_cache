package models

type User struct {
	ID              uint   `json:"id"`
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
}

func NewUser(name, email, password, passwordConfirm string) *User {
	return &User{
		Name:            name,
		Email:           email,
		Password:        password,
		PasswordConfirm: passwordConfirm,
	}
}
