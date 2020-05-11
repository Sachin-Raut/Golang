
/*
write a microservice that receives timesZone abbreviation & returns the current time 
at that timesZone

1. we will add logging middleware
(this will log the time, method & request url)

2. we will also handle errors
*/

package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"
)

var conversionMap = map[string]string {
	"IST":"+5h30m",
	"GMT":"+2h",
	"ART":"-3h",
}

type timeZoneConversion struct {
	TimeZone string
	CurrentTime string
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func main(){
	http.HandleFunc("/convert", loggingMiddleware(handler))

	http.HandleFunc("/", loggingMiddleware(notFoundHandler))
	log.Printf("%s - starting server on port : 8000")
	
	//start the server
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))


	srv := &http.Server {
		Addr : ":8000",
		ReadTimeout : 10 * time.Second,
		WriteTimeout : 10 * time.Second,
		MaxHeaderBytes : 1 << 20,		
	}

	log.Fatal(srv.ListenAndServe())

}

func notFoundHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error 404 - The requested URL doesn't exist")
}

func handler(w http.ResponseWriter, r *http.Request){
	timeZone := r.URL.Query().Get("tz")

	if timeZone == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error 400 - tz query parameter is required")
		return
	}

	timeDifference, ok := conversionMap[timeZone]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,`Error 404 - the tz value doesn't correspond to an existing tz value`, timeZone)
		return
	}


	currentTimeConverted, err := getCurrentTimeByTimeDifference(timeDifference)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error - Server error")
		return
	}

	w.WriteHeader(http.StatusOK)

	tzc := new(timeZoneConversion)
	tzc.CurrentTime = currentTimeConverted
	tzc.TimeZone = timeZone

	jsonResponse, err := json.Marshal(tzc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error - Server error")
		return
	}
	fmt.Fprintf(w, string(jsonResponse))
}

func getCurrentTimeByTimeDifference(timeDifference string)(string, error){
	now := time.Now().UTC()
	difference, err := time.ParseDuration(timeDifference)
	if err != nil {
		return "", err
	}
	now = now.Add(difference)
	return now.Format("15:04:05"),nil
}

func loggingMiddleware(handler handlerFunc) handlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s - %s", time.Now().Format("2018-11-25 14:32:58"), r.Method, r.URL.String())
		handler(w,r)
	}
	return fn
}