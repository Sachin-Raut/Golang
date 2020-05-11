//CRUD without DB

package main

import (
	"log"
	"time"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

//Book is
type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year string `json:"year"`
}

var books []Book

func main(){

	books = append(books,
	Book{ID:1, Title: "book 1", Author: "Author 1", Year: "2010"},
	Book{ID:2, Title: "book 2", Author: "Author 2", Year: "2010"},
	Book{ID:3, Title: "book 3", Author: "Author 3", Year: "2010"},
	Book{ID:4, Title: "book 4", Author: "Author 4", Year: "2010"},
	Book{ID:5, Title: "book 5", Author: "Author 5", Year: "2010"},
)

	router := mux.NewRouter()

	router.HandleFunc("/books",getBooks).Methods("GET")
	router.HandleFunc("/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/books",addBook).Methods("POST")
	router.HandleFunc("/books",updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

	//start the server
	// log.Fatal(http.ListenAndServe(":8000", router))

	srv := &http.Server {
		Handler : router,
		Addr : ":8000",

		//enforce timeouts for servers
		ReadTimeout : 10 * time.Second,
		WriteTimeout : 10 * time.Second,
			
	}

	log.Fatal(srv.ListenAndServe())

}

func getBooks(w http.ResponseWriter, r *http.Request){
	log.Println("Get all books is called")

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	log.Println("getBook")

	params := mux.Vars(r)

	log.Println(params)

	i, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request){
	log.Println("addBook")
	var book Book
	
	_ = json.NewDecoder(r.Body).Decode(&book)

	//return response containing all books

	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request){
	log.Println("updateBook")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}
	//lets return all the books

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request){
	log.Println("removeBook")

	params := mux.Vars(r)

	//converts string to int

	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	//let's return all the books
	json.NewEncoder(w).Encode(books)
}