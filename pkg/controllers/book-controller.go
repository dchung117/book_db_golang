package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dchung117/book_db_golang/pkg/models"
	"github.com/dchung117/book_db_golang/pkg/utils"
	"github.com/gorilla/mux"
)

// create new book variable
var NewBook models.Book

// get all books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	// get slice w/ all books
	allBooks := models.GetAllBooks()

	// convert data to JSON
	res, _ := json.Marshal(allBooks)

	// set header as JSON encoding
	w.Header().Set("Content-Type", "pkglication/json")

	// set header w/ status 200
	w.WriteHeader(http.StatusOK)

	// return JSON of all books
	w.Write(res)
}

// get single book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	// parse request body variables
	params := mux.Vars(r)

	// get book ID
	bookId := params["id"]

	// convert string to int
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error when parsing")
	}

	// get book and db object
	book, _ := models.GetBook(Id)

	// convert data to JSON
	res, _ := json.Marshal(book)

	// set header to JSON encoding
	w.Header().Set("Content-Type", "pkglication/json")

	// set header w/ status 200
	w.WriteHeader(http.StatusOK)

	// write JSON of book details
	w.Write(res)
}

// create book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// create new book interface
	newBook := &models.Book{}

	// parse out book from JSON, store in book interface
	utils.ParseBody(r, newBook)

	// create book in db
	b := newBook.CreateBook()

	// encode book into JSON
	res, _ := json.Marshal(b)

	// set up header
	w.Header().Set("Content-Type", "pkglication")

	// status 200
	w.WriteHeader(http.StatusOK)

	// write JSON of new book
	w.Write(res)

}

// delete book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// parse out query arguments
	params := mux.Vars(r)
	bookId := params["id"]

	// convert id to integer
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error when parsing")
	}
	// delete book from db
	deletedBook := models.DeleteBook(id)

	// json encode book
	res, _ := json.Marshal(deletedBook)

	// set JSON header
	w.Header().Set("Content-Type", "pkglication/json")

	// status 200
	w.WriteHeader(http.StatusOK)

	// write JSON result
	w.Write(res)
}

// updating a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// parse ID from request body
	params := mux.Vars(r)
	bookId := params["id"]

	// convert bookId to integer
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error when parsing")
	}

	// create book
	var newBook = &models.Book{}

	// convert book info from request body from JSON
	utils.ParseBody(r, newBook)

	// get book by the ID, update its data and save to db
	book, db := models.GetBook(id)
	if newBook.Name != "" {
		book.Name = newBook.Name
	}
	if newBook.Author != "" {
		book.Author = newBook.Author
	}
	if newBook.Publication != "" {
		book.Publication = newBook.Publication
	}
	db.Save(&book)

	// convert b to JSON
	res, _ := json.Marshal(book)

	// JSON header
	w.Header().Set("Content-Type", "pkglication/json")

	// status 200
	w.WriteHeader(http.StatusOK)

	// write response
	w.Write(res)
}
