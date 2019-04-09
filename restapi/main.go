package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Init Book var as a slice Book struct
var books []Book
//Author Structure (Model)
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Book Structure (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Auther *Author `json:"auther"`
}

func main() {
	//Init Router
	r := mux.NewRouter()

	// Mock Data @todo -implement DB
	books = append(books, Book{ID: "1", Isbn: "44889", Title: "Book one ", Auther: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "45789", Title: "Book two ", Auther: &Author{FirstName: "Jane", LastName: "Doe"}})

	// Route Handlers /Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/book", createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", updateBook).Methods("UPDATE")
	r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func deleteBook(writer http.ResponseWriter, r *http.Request) {

}

func updateBook(writer http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r) //Get params
	// Loop through books and find with id
	for _,item := range books{
		if item.ID ==params["id"] {
				json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
