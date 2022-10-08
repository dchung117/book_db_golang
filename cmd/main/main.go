package main

import (
	"log"
	"net/http"

	"github.com/dchung117/book_db_golang/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// create a router
	r := mux.NewRouter()

	// register bookstore routes
	routes.RegisterBookDBRoutes(r)

	// set router as defaultservemux
	http.Handle("/", r)

	// launch router, log error if fails
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
