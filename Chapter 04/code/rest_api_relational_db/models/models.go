package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Employee struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

var DB *gorm.DB

func ConnectDatabase() {

	database, error := gorm.Open(sqlite.Open("customers.db"), &gorm.Config{})

	if error != nil {
		panic("Failed to connect to database!")
	}

	error = database.AutoMigrate(&Employee{})
	if error != nil {
		return
	}

	DB = database
}
