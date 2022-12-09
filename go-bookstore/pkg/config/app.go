package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	dsn := `root:Kartik34@tcp(localhost:3306)/go-bookstore?charset=utf8&parseTime=True&loc=Local`
	d, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
