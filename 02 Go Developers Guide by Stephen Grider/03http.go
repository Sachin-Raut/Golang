package main

import (
	"os"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error -", err)
		os.Exit(1)
	}
	fmt.Println(resp) //in "resp", we will not see body

	byteSlice := make([]byte, 99999)

	/*
	1. this byteSlice has been initialised with 99999 empty elements
	2. make() takes first argument as slice of any type (in this case byte)
	3. we initialised byte slice with giant capacity because Read function 
	puts data & quits as soon as the byteSlice is full.
	4. 99999 is big enough for all the body data to fit into byteSlice
	*/

	resp.Body.Read(byteSlice)
	fmt.Println(string(byteSlice))
}