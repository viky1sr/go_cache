package providers

import (
	"database/sql"

	"github.com/viky1sr/go_cache.git/app/repositories"
	"github.com/viky1sr/go_cache.git/app/services"
)

// AuthProvider is responsible for providing the necessary dependencies for the auth module
type AuthProvider struct{}

// ProvideAuthRepository provides the auth repository instance
func (provider *AuthProvider) ProvideAuthRepository(db *sql.DB) *repositories.SQLAuthRepository {
	return repositories.NewSQLAuthRepository(db)
}

// ProvideAuthService provides the auth service instance
func (provider *AuthProvider) ProvideAuthService(userRepo repositories.UserRepository) services.AuthService {
	return services.NewAuthService(userRepo)
}
