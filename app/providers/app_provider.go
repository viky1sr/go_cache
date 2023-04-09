package providers

import (
	"database/sql"
	"fmt"
	"github.com/viky1sr/go_cache.git/app/validators"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"github.com/viky1sr/go_cache.git/app/controllers"
	"github.com/viky1sr/go_cache.git/app/repositories"
	"github.com/viky1sr/go_cache.git/app/services"
	"github.com/viky1sr/go_cache.git/config"
)

// AppProvider is responsible for providing the necessary dependencies for the app
type AppProvider struct {
	Host string
	Port string
}

func (a *AppProvider) GetHost() string {
	if a.Host == "" {
		a.Host = "localhost"
	}
	return a.Host
}

func (a *AppProvider) SetHost(host string) {
	a.Host = host
}

func (a *AppProvider) GetPort() string {
	if a.Port == "" {
		a.Port = "8080"
	}
	return a.Port
}

func (a *AppProvider) SetPort(port string) {
	a.Port = port
}

// ProvideRouter provides the router instance for the app
func (provider *AppProvider) ProvideRouter() *mux.Router {
	return mux.NewRouter()
}

// ProvideDB provides the database instance for the app
func (provider *AppProvider) ProvideDB() (*sql.DB, error) {
	dbConfig := config.GetDBConfig()

	connectionString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	)

	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ProvideBookController provides the book controller instance for the app
func (provider *AppProvider) ProvideBookController() *controllers.BookController {
	db, err := provider.ProvideDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s", err)
	}

	bookRepository := repositories.NewSqlBookRepository(db)
	bookValidator := validators.NewBookValidator()
	bookService := services.NewBookService(bookRepository, bookValidator)

	return controllers.NewBookController(bookService)
}

// ProvideUserController provides the user controller instance for the app
func (provider *AppProvider) ProvideUserController() *controllers.UserController {
	db, err := provider.ProvideDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s", err)
	}

	userRepository := repositories.NewSqlUserRepository(db)
	bookValidator := validators.NewUserValidator()
	userService := services.NewUserService(userRepository, bookValidator)

	return controllers.NewUserController(userService)
}
