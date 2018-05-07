package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Book struct {
	ID     string `json:"id"`
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

func buildList() []Book {
	author1 := Author{"Nuttapong", "Rojanavanich"}
	book1 := Book{"1", "77777", "How to play game", author1}
	author2 := Author{"Anonymous", "Uncle"}
	book2 := Book{"2", "12345", "How to sleep", author2}
	author3 := Author{"Hunter", "Jagras"}
	book3 := Book{"3", "93867", "How to cook Well Done Steak", author3}
	book4 := Book{"4", "67890", "How to drink Potion like a pro", author3}

	books = append(books, book1)
	books = append(books, book2)
	books = append(books, book3)
	books = append(books, book4)
	return books
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(50000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}

var books []Book

func main() {
	books = buildList()

	r := mux.NewRouter()
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createNewBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":4545", r))
}
