package routes

import (
	"github.com/dchung117/book_db_golang/pkg/controllers"
	"github.com/gorilla/mux"
)

// declare routing function handle
var RegisterBookDBRoutes = func(router *mux.Router) {
	// handles
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
