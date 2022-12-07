package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.getBookById).Methods("GET")
	router.HandleFunc("/books", controllers.createBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.deleteBook).Methods("DELETE")
}
