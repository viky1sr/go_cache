package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viky1sr/go_cache.git/app/providers"
)

func RegisterAuthRoutes(router *mux.Router, provider *providers.AppProvider) {
	authController := provider.ProvideAuthController()

	router.HandleFunc("/login", authController.Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", authController.Logout).Methods(http.MethodPost)
}
