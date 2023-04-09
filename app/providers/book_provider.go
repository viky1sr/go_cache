package providers

import (
	"database/sql"
	"github.com/viky1sr/go_cache.git/app/validators"

	"github.com/viky1sr/go_cache.git/app/repositories"
	"github.com/viky1sr/go_cache.git/app/services"
)

// BookProvider is responsible for providing the necessary dependencies for the book module
type BookProvider struct{}

// ProvideBookRepository provides the book repository instance
func (provider *BookProvider) ProvideBookRepository(db *sql.DB) repositories.BookRepository {
	return repositories.NewSqlBookRepository(db)
}

// ProvideBookService provides the book service instance
func (provider *BookProvider) ProvideBookService(repo repositories.BookRepository, validator *validators.BookValidator) services.BookService {
	return services.NewBookService(repo, validator)
}
