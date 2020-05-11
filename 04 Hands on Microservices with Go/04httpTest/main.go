package main

import (
	"fmt"
	"net/http"
	"hands-on-microservices/04httptest/handlers"
)

func main(){
	fmt.Println("hi")
	http.HandleFunc("/example",handlers.MyHandler)

	// http.ListenAndServe(":8080", nil)
	// fmt.Println("hi1")

	srv := &http.Server {
		Addr : ":8080",
		ReadTimeout : 10 * time.Second,
		WriteTimeout : 10 * time.Second,
		MaxHeaderBytes : 1 << 20,		
	}

	log.Fatal(srv.ListenAndServe())
}