package models

import (
	_ "example.com/my_mysql_driver"
	"github.com/go-delve/delve/pkg/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Crete(&b)

	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)

	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)

	return &book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)

	return book
}
