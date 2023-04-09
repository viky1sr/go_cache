package repositories

import (
	"github.com/viky1sr/go_cache.git/app/models"
)

type BookRepository interface {
	GetAllBooks() ([]*models.Book, error)
	GetBookByID(id int) (*models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(id int, book *models.Book) error
	DeleteBook(id int) error
}
