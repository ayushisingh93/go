package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct
type Books struct {
	Id     string  `json:id`
	Isbn   string  `json:isbn`
	Title  string  `json:title`
	Author *Author `json:author`
}

// Author struct
type Author struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

// silce
var books []Books

// Get all book
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get all book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Get all book
func createBooks(w http.ResponseWriter, r *http.Request) {

}

// Get all book
func updateBooks(w http.ResponseWriter, r *http.Request) {

}

// Get all book
func deleteBooks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//init router
	route := mux.NewRouter()

	//Moke Data
	books = append(books, Books{Id: "1", Isbn: "12345", Title: "Book one", Author: &Author{Firstname: "Ritesh", Lastname: "Kumar"}})
	books = append(books, Books{Id: "2", Isbn: "12345", Title: "Book two", Author: &Author{Firstname: "Ayushi", Lastname: "Singh"}})

	//Router handler Endpoints
	route.HandleFunc("/api/books", getBooks).Method("GET")
	route.HandleFunc("/api/books/{id}", getBook).Method("GET")
	route.HandleFunc("/api/books", createBooks).Method("POST")
	route.HandleFunc("/api/books/{id}", updateBooks).Method("PUT")
	route.HandleFunc("/api/books/{id}", deleteBooks).Method("DELETE")
	log.Fatal(http.ListenAndServe(":8000", route))
}
