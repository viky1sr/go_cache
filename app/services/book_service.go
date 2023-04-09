package services

import (
	"errors"
	"github.com/viky1sr/go_cache.git/app/models"
	"github.com/viky1sr/go_cache.git/app/repositories"
	"github.com/viky1sr/go_cache.git/app/validators"
)

type BookService interface {
	GetAllBooks() ([]*models.Book, error)
	GetBookByID(id int) (*models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(id int, book *models.Book) error
	DeleteBook(id int) error
}

type bookService struct {
	bookRepo      repositories.BookRepository
	bookValidator *validators.BookValidator
}

func NewBookService(bookRepo repositories.BookRepository, bookValidator *validators.BookValidator) BookService {
	return &bookService{bookRepo: bookRepo, bookValidator: bookValidator}
}

func (s *bookService) GetAllBooks() ([]*models.Book, error) {
	return s.bookRepo.GetAllBooks()
}

func (s *bookService) GetBookByID(id int) (*models.Book, error) {
	if id <= 0 {
		return nil, errors.New("invalid book ID")
	}

	return s.bookRepo.GetBookByID(id)
}

func (s *bookService) CreateBook(book *models.Book) error {
	err := s.bookValidator.Validate(book)
	if err != nil {
		return err
	}
	return s.bookRepo.CreateBook(book)
}

func (s *bookService) UpdateBook(id int, book *models.Book) error {
	if id <= 0 {
		return errors.New("invalid book ID")
	}

	if book == nil {
		return errors.New("book cannot be nil")
	}

	return s.bookRepo.UpdateBook(id, book)
}

func (s *bookService) DeleteBook(id int) error {
	if id <= 0 {
		return errors.New("invalid book ID")
	}

	return s.bookRepo.DeleteBook(id)
}
