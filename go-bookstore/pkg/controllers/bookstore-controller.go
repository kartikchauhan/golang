package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kartikchauhan/go-bookstore/pkg/models"
	"github.com/kartikchauhan/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	utils.PrintMethod("GetBooks")

	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	utils.PrintMethod("GetBookById")

	params := mux.Vars(r)
	id := params["id"]

	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book, _ := models.GetBookById(bookId)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	utils.PrintMethod("CreateBook")

	Book := &models.Book{}
	utils.ParseBody(r, Book)

	book := Book.CreateBook()

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	utils.PrintMethod("DeleteBook")

	params := mux.Vars(r)
	id := params["id"]

	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}

	book := models.DeleteBook(bookId)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	utils.PrintMethod("UpdateBook")

	updatedBookDetails := &models.Book{}
	utils.ParseBody(r, updatedBookDetails)

	params := mux.Vars(r)
	id := params["id"]

	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}

	bookDetails, db := models.GetBookById(bookId)

	if updatedBookDetails.Name != "" {
		bookDetails.Name = updatedBookDetails.Name
	}

	if updatedBookDetails.Author != "" {
		bookDetails.Author = updatedBookDetails.Author
	}

	if updatedBookDetails.Publication != "" {
		bookDetails.Publication = updatedBookDetails.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
