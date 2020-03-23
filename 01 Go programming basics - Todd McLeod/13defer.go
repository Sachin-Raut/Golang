package main

import "fmt"

func main(){
	i := 0
	defer fmt.Println(i) 
	i = 2
	return
}

//output is 0 because during defer call i was 0