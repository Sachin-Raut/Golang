package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(){
		ch <- 10
		close(ch)
	}()

	v, ok := <- ch
	fmt.Println(v, ok) //10, true
	fmt.Println(<-ch)  //0 because we are retrieving value after the channel is closed, otherwise we would have got deadlock error
}