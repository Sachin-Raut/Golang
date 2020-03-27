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
	
	//unbuffered channel
	ch := make(chan int)  //bidirectional channel
	go func(){
		ch <- 20          //sending value to channel (send channel)
	}()
	fmt.Println(<- ch)   //receiving value from channel (receive channel)


	//unbuffered channel
	message := make(chan string) //bidirectional channel
	go func(){
		time.Sleep(time.Second)  
		message <- "hi"         //sending value to channel (send channel)
	}()
	fmt.Println(<- message)     //receiving value from channel (receive channel)


	//buffered channel. Here u can send only 2 values
	c := make(chan int, 2)   //bidirectional channel
	c <- 15					 //sending value to channel (send channel)
	fmt.Println(<-c)		 //receiving value from channel (receive channel)

	c <- 200
	fmt.Println(<-c)
}