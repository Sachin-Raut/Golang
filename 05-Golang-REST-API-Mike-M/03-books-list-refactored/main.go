package main

import (
	
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"net/http"
	// "encoding/json"
	"github.com/gorilla/mux"
	"05-Golang-REST-API-Mike-M/03-books-list-refactored/driver"
	"05-Golang-REST-API-Mike-M/03-books-list-refactored/controllers"
)


func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}


// var books []Book


func main(){

	db := driver.ConnectDB()

	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Starting the server1")
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

/*

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

*/