package routes

import (
	"github.com/gorilla/mux"
	"github.com/viky1sr/go_cache.git/app/providers"
)

// RegisterRoutes registers all the routes
func RegisterRoutes(router *mux.Router, appProvider *providers.AppProvider) {
	authRouter := router.PathPrefix("/api").Subrouter()
	RegisterAuthRoutes(authRouter, appProvider)
	RegisterUserRoutes(authRouter, appProvider)
	RegisterBookRoutes(authRouter, appProvider)
}
