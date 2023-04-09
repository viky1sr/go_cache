package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viky1sr/go_cache.git/app/providers"
)

// RegisterBookRoutes registers book related routes
func RegisterBookRoutes(router *mux.Router, provider *providers.AppProvider) {
	bookController := provider.ProvideBookController()

	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			bookController.GetAllBooks(w, r)
		case "POST":
			bookController.CreateBook(w, r)
		default:
			http.NotFound(w, r)
		}
	}).Methods("GET", "POST")

	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			bookController.GetBookByID(w, r)
		case "PUT":
			bookController.UpdateBook(w, r)
		case "DELETE":
			bookController.DeleteBook(w, r)
		default:
			http.NotFound(w, r)
		}
	}).Methods("GET", "PUT", "DELETE")
}
