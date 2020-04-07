package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",handler)

	//start the server
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World Sachin")
}