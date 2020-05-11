//CRUD with DB

package main

import (

	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"time"
	// "strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)


func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}


// Book is
type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year string `json:"year"`
}

var books []Book

const (
	//sslmode can be disable or verify-full
	host     = "satao.db.elephantsql.com"
	port     = 5432
	user     = "bgligpzc"
	password = "cm_uF7PbzYycl39koGWlaCLAEd6mnuF9"
	dbname   = "bgligpzc"
	sslmode  = "verify-full"
	// sslmode  = "disable" 
  )


var db *sql.DB
var err error

func main(){

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=%s",
	host, port, user, password, dbname, sslmode)
	
	db, err = sql.Open("postgres", connStr)

	logFatal(err)

	db.Ping()

	router := mux.NewRouter()

	router.HandleFunc("/books",getBooks).Methods("GET")
	router.HandleFunc("/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/books",addBook).Methods("POST")
	router.HandleFunc("/books",updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

	fmt.Println("Starting the server")

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

	var book Book

	books = []Book{}

	rows, err := db.Query("select * from booksTable")

	logFatal(err)
	defer rows.Close()

	for rows.Next(){
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}


func getBook(w http.ResponseWriter, r *http.Request){
	log.Println("getBook")

	var book Book
	params := mux.Vars(r)

	rows := db.QueryRow("select * from booksTable where id=$1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	logFatal(err)

	json.NewEncoder(w).Encode(book)
}


func addBook(w http.ResponseWriter, r *http.Request){
	log.Println("addBook")
	var book Book
	var bookID int

	_ = json.NewDecoder(r.Body).Decode(&book)

	err = db.QueryRow("insert into booksTable (title, author, year) values ($1,$2,$3) returning id;", book.Title, book.Author, book.Year).Scan(&bookID)

	logFatal(err)

	//return response containing all books

	json.NewEncoder(w).Encode(bookID)
}


func updateBook(w http.ResponseWriter, r *http.Request){
	log.Println("updateBook")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	result, err := db.Exec("update booksTable set title = $1, author = $2, year = $3 where id = $4 returning id", &book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	//lets return the number of rowsAffected
	json.NewEncoder(w).Encode(rowsUpdated)
}


func removeBook(w http.ResponseWriter, r *http.Request){
	log.Println("removeBook")

	params := mux.Vars(r)

	result, err := db.Exec("delete from booksTable where id = $1", params["id"])

	logFatal(err)

	rowsDeleted, err := result.RowsAffected()

	logFatal(err)

	//let's return the number of deleted
	json.NewEncoder(w).Encode(rowsDeleted)
}

