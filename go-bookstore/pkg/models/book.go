package models

import (
	_ "example.com/my_mysql_driver"
	"github.com/go-delve/delve/pkg/config"
	"gorm.io/gorm"
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
