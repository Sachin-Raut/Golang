package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"06-GOLANG-REST-JWT-Mike-M/02-JWT-With-refactor/driver"	
	"06-GOLANG-REST-JWT-Mike-M/02-JWT-With-refactor/utils"
	"06-GOLANG-REST-JWT-Mike-M/02-JWT-With-refactor/controllers"
)

var db *sql.DB
var err error

func main() {

	db = driver.ConnectDB()

	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")

	router.HandleFunc("/login", controller.Login(db)).Methods("POST")

	router.HandleFunc("/protected", utils.TokenVerify(controller.ProtectedEndpoint())).Methods("GET")

	//start server
	fmt.Println("Started server123")
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



