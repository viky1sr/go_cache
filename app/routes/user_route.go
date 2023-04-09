package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viky1sr/go_cache.git/app/providers"
)

// RegisterUserRoutes registers user related routes
func RegisterUserRoutes(router *mux.Router, provider *providers.AppProvider) {
	userController := provider.ProvideUserController()

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			userController.GetAllUsers(w, r)
		case "POST":
			userController.CreateUser(w, r)
		default:
			http.NotFound(w, r)
		}
	}).Methods("GET", "POST")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			userController.GetUserByID(w, r)
		case "PUT":
			userController.UpdateUser(w, r)
		case "DELETE":
			userController.DeleteUser(w, r)
		default:
			http.NotFound(w, r)
		}
	}).Methods("GET", "PUT", "DELETE")
}
