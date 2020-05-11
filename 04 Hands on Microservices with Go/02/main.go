
/*
write a microservice that receives timesZone abbreviation & returns the current time 
at that timesZone
*/

package main

import (
	"fmt"
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
	timeZone := r.URL.Query().Get("tz")
	timeDifference, _ := conversionMap[timeZone]
	currentTimeConverted, _ := getCurrentTimeByTimeDifference(timeDifference)

	tzc := new(timeZoneConversion)
	tzc.CurrentTime = currentTimeConverted
	tzc.TimeZone = timeZone

	jsonResponse, _ := json.Marshal(tzc)
	w.WriteHeader(http.StatusOK)
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