package controllers

import (
	"apiv1/data"
	"apiv1/datatypes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// ...
	books := data.Books
	var listOfBooks datatypes.Books

	for id, name := range books {
		book := datatypes.Book{Id: id, Name: name}
		listOfBooks.AddBook(book)
	}

	jsonBooks, err := json.Marshal(listOfBooks)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBooks)
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	// ...
	bookIdStr := mux.Vars(r)["id"]
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad request"))
		return
	}
	book, ok := data.Books[bookId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(book))

	marshed, err := json.Marshal(book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(marshed)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// ...
	var newBooks datatypes.Book

	err := json.NewDecoder(r.Body).Decode(&newBooks)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	_, ok := data.Books[newBooks.Id]
	if ok {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Book already exists"))
		return
	}

	data.Books[newBooks.Id] = newBooks.Name
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book created"))

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// ...

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// ...
	bookIdStr := mux.Vars(r)["id"]
	bookId, err := strconv.Atoi(bookIdStr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad request"))
		return
	}

	_, ok := data.Books[bookId]
	if ok {
		delete(data.Books, bookId)
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Book deleted"))
		return
	}
}
