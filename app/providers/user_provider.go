package providers

import (
	"database/sql"
	"github.com/viky1sr/go_cache.git/app/validators"

	"github.com/viky1sr/go_cache.git/app/repositories"
	"github.com/viky1sr/go_cache.git/app/services"
)

// UserProvider is responsible for providing the necessary dependencies for the user module
type UserProvider struct{}

// ProvideUserRepository provides the user repository instance
func (provider *UserProvider) ProvideUserRepository(db *sql.DB) repositories.UserRepository {
	return repositories.NewSqlUserRepository(db)
}

// ProvideUserService provides the user service instance
func (provider *UserProvider) ProvideUserService(repo repositories.UserRepository, validator *validators.UserValidator) services.UserService {
	return services.NewUserService(repo, validator)
}
