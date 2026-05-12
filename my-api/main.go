package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var books = []Book{
	{ID: 1, Title: "The Go Programming Language", Author: "Donovan & Kernighan"},
	{ID: 2, Title: "Clean Code", Author: "Robert C. Martin"},
}

var validate = validator.New()

func sendJSON(w http.ResponseWriter, status int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, http.StatusOK, APIResponse{
		Status: "success", Message: "Books retrieved", Data: books,
	})
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		sendJSON(w, http.StatusBadRequest, APIResponse{
			Status: "error", Message: "Invalid JSON format",
		})
		return
	}
	if err := validate.Struct(newBook); err != nil {
		sendJSON(w, http.StatusBadRequest, APIResponse{
			Status: "error", Message: err.Error(),
		})
		return
	}
	newBook.ID = len(books) + 1
	books = append(books, newBook)
	sendJSON(w, http.StatusCreated, APIResponse{
		Status: "success", Message: "Book created", Data: newBook,
	})
}

func main() {
	// Go 1.22+ method routing — method salah = otomatis 405
	http.HandleFunc("GET /books", getBooks)
	http.HandleFunc("POST /books", addBook)

	fmt.Println("Server running di :8080")
	http.ListenAndServe(":8080", nil)
}
