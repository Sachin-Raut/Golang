//context is used to cancel a running goroutine

package main

import (
	"fmt"
	"time"
	"context"
	"math/rand"
)

func main() {
	//timeout is 80 millisecond
	duration := 80 * time.Millisecond

	//context with timeout of 80 millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)

	defer cancel()

	ch := make(chan string, 1)

	go func(){

		sleepTime := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(sleepTime)

		//send value to a channel
		ch <- "paper"
	}()

	select {
	case p := <- ch :
		fmt.Println("work complete",p)
	case <- ctx.Done():
		fmt.Println("waited for 80 milliseconds, now moving on")
	}
}

/*

case p is executed if the channel receives value within 80 milliseconds
ctx.Done() is executed if the channel doesn't receive value within 80 milliseconds

*/