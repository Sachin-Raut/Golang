/*

Drop pattern

*/

package main

import "fmt"

func main(){

	//buffered channel of capacity 5
	const cap = 5 

	ch := make(chan string, cap)

	go func(){
		for p := range ch { //this is blocking call
			fmt.Println("emp received -", p)
		}
	}()

	const work = 10

	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager - send ack")
		default:
			fmt.Println("manager - drop")
		}
	}
	close(ch)
}

/*

1. we are creating buffered channel of capacity 5
2. "for p := range ch" this is blocking call
3. we are sending 10 values, but the buffer capacity of the channel is 5. 
So it may or may not receive 5 values & other values may be dropped.

A buffered channel greater than 1 provides no guarantee that a signal being 
sent is ever received.

*/
