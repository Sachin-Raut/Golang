package controllers

import (
	"log"
	"database/sql"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	"05-Golang-REST-API-Mike-M/03-books-list-refactored/model"
	"05-Golang-REST-API-Mike-M/03-books-list-refactored/utils"
	bookrepository "05-Golang-REST-API-Mike-M/03-books-list-refactored/repository/b"
	

)


func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}

//Controller is
type Controller struct {}

var books []model.Book

//GetBooks is
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		log.Println("Get all books is called")
	
		var book model.Book
		var error model.Error
	
		books = []model.Book{}

		bookRepo := bookrepository.BookRepository{}

		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type","application/json")
		utils.SendSuccess(w,books)
	}	
}


//GetBook is
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		log.Println("getBook")
	
		var book model.Book
		var error model.Error

		params := mux.Vars(r)
	
		//convert string to int

		id, _ := strconv.Atoi(params["id"])
		
		bookRepo := bookrepository.BookRepository{}

		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not found"
				utils.SendError(w, http.StatusNotFound, error)
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
			}
			return
		}
		w.Header().Set("Content-Type","application/json")
		utils.SendSuccess(w, book)
	}
}

//AddBook is
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log.Println("addBook")
		var book model.Book
		var error model.Error
	
		_ = json.NewDecoder(r.Body).Decode(&book)

		log.Println(book.ID)
		log.Println(book.Title)
		log.Println(book.Author)
		log.Println(book.Year)
	
		if book.Title == "" || book.Author == "" || book.Year == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookrepository.BookRepository{}

		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type","text/plain")
		utils.SendSuccess(w, bookID)
	}
}

//UpdateBook is
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		log.Println("updateBook")
	
		var book model.Book
		var error model.Error
	
		_ = json.NewDecoder(r.Body).Decode(&book)

		log.Println(book.ID)
		log.Println(book.Title)
		log.Println(book.Author)
		log.Println(book.Year)
	
		if book.ID == 0 || book.Title == "" || book.Author == "" || book.Year == "" {
			error.Message = "All fields are required...."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookrepository.BookRepository{}

		updatedRows, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "All fields are required"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, updatedRows)
	}	
}

//RemoveBook is 
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

		log.Println("removeBook")
	
		var error model.Error
		params := mux.Vars(r)

		//convert string to int

		id, _ := strconv.Atoi(params["id"])

		bookRepo := bookrepository.BookRepository{}

		rowsDeleted, err := bookRepo.RemoveBook(db, id)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsDeleted == 0 {
			error.Message = "Not found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}
		w.Header().Set("Content-Type","text/plain")
		utils.SendSuccess(w,rowsDeleted)
	}
}
