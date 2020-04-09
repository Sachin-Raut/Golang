package main

import (
	"fmt"
	"net/http"
	"hands-on-microservices/04httptest/handlers"
)

func main(){
	fmt.Println("hi")
	http.HandleFunc("/example",handlers.MyHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("hi1")
}