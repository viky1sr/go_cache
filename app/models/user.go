package models

type User struct {
	ID              uint   `json:"id"`
	Name            string `json:"name" validate:"required" error:"Name is required"`
	Email           string `json:"email" validate:"required,email" error:"Email is required and must be in valid email format"`
	Password        string `json:"password" validate:"required,min=6" error:"Password is required and must be at least 6 characters long"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password" error:"Password confirmation must match with password"`
}

func NewUser(name, email, password, passwordConfirm string) *User {
	return &User{
		Name:            name,
		Email:           email,
		Password:        password,
		PasswordConfirm: passwordConfirm,
	}
}
