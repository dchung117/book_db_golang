package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // only to initialize side-effects, not for using objects
)

// create database
var (
	db *gorm.DB
)

// connect to db
func Connect() {
	// open connection
	d, err := gorm.Open("mysql", "local_user:Test123!@/books_db?charset=utf8&parseTime=True&loc=Local")

	// throw error
	if err != nil {
		panic(err)
	}

	// otherwise, assign global variable db to connection
	db = d
}

// return db
func GetDB() *gorm.DB {
	return db
}
