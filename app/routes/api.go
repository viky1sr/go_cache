package routes

import (
	"github.com/gorilla/mux"
	"github.com/viky1sr/go_cache.git/app/providers"
)

// RegisterRoutes registers all the routes
func RegisterRoutes(router *mux.Router, appProvider *providers.AppProvider) {
	RegisterAuthRoutes(router, appProvider)
	RegisterUserRoutes(router, appProvider)
	RegisterBookRoutes(router, appProvider)
}
