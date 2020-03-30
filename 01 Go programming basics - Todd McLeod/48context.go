package main

import (
	"fmt"
	"time"
)

func main(){
	c := generate()
	receive(c)
	fmt.Println("about to exit")
}

func generate() <- chan int {
	c := make(chan int)

	go func(){
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			c <- i //send channel
			fmt.Println("Send")
		}
		close(c)
	}()
	return c
}

func receive(c <- chan int) {
	for v := range c { //blocking call 
		fmt.Println(v)
	}
}

/*

1. as soon as value is sent over channel in "generate function", it is returned
2. that returned value is sent to "receive func" & it is printed.

in nutshell, as soon as value is sent, its printed.

*/