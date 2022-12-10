package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kartikchauhan/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Println("Routes registered")

	log.Println("Starting server on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Unable to start server")
	}
}
