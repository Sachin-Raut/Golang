package main

import (
	"fmt"
	// "sync"
	"time"
)

func main(){
	ch := make(chan string)


	go func(){
		time.Sleep(time.Second * 5)
		ch <- "sachin"
	}()

	fmt.Println(<-ch)
}