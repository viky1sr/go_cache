package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/viky1sr/go_cache.git/app/models"
)

type BookValidator struct {
	validate *validator.Validate
}

func NewBookValidator() *BookValidator {
	validate := validator.New()
	return &BookValidator{validate: validate}
}

func (v *BookValidator) Validate(book *models.Book) error {
	return v.validate.Struct(book)
}
