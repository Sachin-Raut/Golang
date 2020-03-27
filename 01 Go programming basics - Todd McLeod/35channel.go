/*

1. channels are reference types
2. channels are better way for synchronising code & writing concurrent code
3. channels are pipes that connect concurrent goroutines
4. you can send values from 1 goroutine, & receive that value in another goroutine
5. by default, channels are "unbuffered", meaning that they will only accept send (chan <- )
   if there is a corresponding receive (<- chan)

*/

package main 

import (
	"fmt"
	"time"
)

func main() {
	
	ch := make(chan int)

	go func(){
		ch <- 20
	}()

	fmt.Println(<- ch)

	message := make(chan string)

	go func(){
		time.Sleep(time.Second)
		message <- "hi"
	}()

	fmt.Println(<- message)
}