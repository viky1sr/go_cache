package routes

import (
	"github.com/viky1sr/go_cache.git/app/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viky1sr/go_cache.git/app/providers"
)

// RegisterBookRoutes registers book related routes
func RegisterBookRoutes(router *mux.Router, provider *providers.AppProvider) {
	bookController := provider.ProvideBookController()

	jwtMiddleware := middleware.JWTMiddleware

	// Protected routes
	router.Handle("/books", jwtMiddleware(http.HandlerFunc(bookController.GetAllBooks))).Methods("GET")
	router.Handle("/books", jwtMiddleware(http.HandlerFunc(bookController.CreateBook))).Methods("POST")
	router.Handle("/books/{id}", jwtMiddleware(http.HandlerFunc(bookController.GetBookByID))).Methods("GET")
	router.Handle("/books/{id}", jwtMiddleware(http.HandlerFunc(bookController.UpdateBook))).Methods("PUT")
	router.Handle("/books/{id}", jwtMiddleware(http.HandlerFunc(bookController.DeleteBook))).Methods("DELETE")
}
