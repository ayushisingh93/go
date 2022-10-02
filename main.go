package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	//get params in the current request
	params := mux.Vars(r)
	//Setting the type of the writter
	w.Header().Set("Contact-Type", "application/json")
	//loop through the books when to get the match
	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Books{})
}

// Get all book
func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "application/json")
	var book Books
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Int())
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Get all book
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		item.Id = params["id"]
		books = append(books[:index], books[index+1:]...)
		var book Books
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.Id = params["id"]
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		return
	}
	json.NewEncoder(w).Encode(books)
}

// Get all book
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		item.Id = params["id"]
		books = append(books[:index], books[index+1:]...)
		break
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//init router
	route := mux.NewRouter()

	//Moke Data
	books = append(books, Books{Id: "1", Isbn: "12345", Title: "Book one", Author: &Author{Firstname: "Ritesh", Lastname: "Kumar"}})
	books = append(books, Books{Id: "2", Isbn: "12345", Title: "Book two", Author: &Author{Firstname: "Ayushi", Lastname: "Singh"}})

	//Router handler Endpoints
	route.HandleFunc("/api/books", getBooks).Methods("GET")
	route.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	route.HandleFunc("/api/books", createBooks).Methods("POST")
	route.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	route.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", route))
}
