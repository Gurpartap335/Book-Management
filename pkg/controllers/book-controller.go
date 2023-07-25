package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gurpartap335/Book-Management/pkg/models"
	"github.com/Gurpartap335/Book-Management/pkg/utils"
	"github.com/Gurpartap335/go-bookstore/pkg/models"
	"github.com/gorilla/mux"
)

var NewBook model.Book 


func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	var := mux.Vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book()
	utils.Parseody(r, CreateBook)
	b := CreateBook.CreateBook();
	res, _ json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.vars(r)
	bookId := vars["bookId"]
	ID, err := strconv
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}


func UpdateBook(w http.ResponseWriter, r * http.Request) {
	var UpdateBook = &models.Book{}
	utils.Parse(r, updateBook)
	vars := mux.Vars(r)
	bookId : vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booksDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		booksDetails.Author = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		booksDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		booksDetails.Publication = UpdateBook.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}