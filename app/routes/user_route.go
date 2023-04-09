package routes

import (
	"github.com/gorilla/mux"
	"github.com/viky1sr/go_cache.git/app/middleware"
	"github.com/viky1sr/go_cache.git/app/providers"
	"net/http"
)

// RegisterUserRoutes registers user related routes
func RegisterUserRoutes(router *mux.Router, provider *providers.AppProvider) {
	userController := provider.ProvideUserController()

	// Define the JWT middleware
	jwtMiddleware := middleware.JWTMiddleware

	// Protected routes
	router.Handle("/users", jwtMiddleware(http.HandlerFunc(userController.GetAllUsers))).Methods("GET")
	router.Handle("/users", jwtMiddleware(http.HandlerFunc(userController.CreateUser))).Methods("POST")
	router.Handle("/users/{id}", jwtMiddleware(http.HandlerFunc(userController.GetUserByID))).Methods("GET")
	router.Handle("/users/{id}", jwtMiddleware(http.HandlerFunc(userController.UpdateUser))).Methods("PUT")
	router.Handle("/users/{id}", jwtMiddleware(http.HandlerFunc(userController.DeleteUser))).Methods("DELETE")

}
