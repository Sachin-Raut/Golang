package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	http.HandleFunc("/",handler)

	//start the server
	// http.ListenAndServe("localhost:8000", nil)

	srv := &http.Server {
		Addr : ":8000",
		ReadTimeout : 10 * time.Second,
		WriteTimeout : 10 * time.Second,
		MaxHeaderBytes : 1 << 20,		
	}

	srv.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World Sachin")
}