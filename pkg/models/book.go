package models

import (
	"github.com/dchung117/book_db_golang/pkg/config"
	"github.com/jinzhu/gorm"
)

// create db variable
var db *gorm.DB

// create Book model
type Book struct {
	gorm.Model
	Name        string `gorm:"json":"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// initialize db - runs before anything else
func init() {
	// connect to DB
	config.Connect()

	// get DB
	db = config.GetDB()

	// auto-migrate w/ an empty book
	db.AutoMigrate(&Book{})
}

// create a book
func (b *Book) CreateBook() *Book {
	// create a new record
	db.NewRecord(b)

	// create entry in db, store in b variable
	db.Create(&b)

	return b
}

// get all books; return slice of books
func GetAllBooks() []Book {
	// declare slice of books
	var Books []Book

	// get all book records, store in books
	db.Find(&Books)

	return Books
}

// get book by ID
func GetBook(Id int64) (*Book, *gorm.DB) {
	// declare retrieved book
	var book Book

	// find book by id, store in book
	db := db.Where("ID=?", Id).Find(&book)

	return &book, db
}

// delete book by ID
func DeleteBook(Id int64) Book {
	var book Book

	// find book by ID, delete and store contents in book
	db.Where("ID=?", Id).Delete(book)

	return book
}
