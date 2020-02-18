package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Books Struct (Model)
type Book struct {
	ID string  `json:"id"`
	Isbn string  `json:"isbn"`
	Title string  `json:"title"`
	Author *Author  `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//Init books var as a slice Book struct
var books []Book

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}


func createBook(w http.ResponseWriter, r *http.Request)  {

}


func updateBook(w http.ResponseWriter, r *http.Request)  {

}


func deleteBook(w http.ResponseWriter, r *http.Request)  {

}


func main()  {
	//init Router
	r := mux.NewRouter()

	//Mock data - @todo  implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Bad Ass", Author: &Author{Firstname: "Jhon", Lastname: "Snow"}})
	books = append(books, Book{ID: "2", Isbn: "448744", Title: "Capitan America", Author: &Author{Firstname: "Ron", Lastname: "Dallas"}})
	books = append(books, Book{ID: "3", Isbn: "448745", Title: "Go land", Author: &Author{Firstname: "Rob", Lastname: "Pike"}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/book",createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}",deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

