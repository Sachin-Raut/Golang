/*

Buffered channel with capacity 1

Use this when its necessary to know if the previous data that was sent 
has been received before sending the new data

*/

package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	go func(){
		//this is blocking call
		for p := range ch {
			fmt.Println("emp working -", p)
		}
	}()

	const work = 10

	for w := 0; w < work; w++ {
		ch <- "paper"
	}
	close(ch)
}

/*

1. we are creating channel with buffer capacity 1
2. "for p := range ch" its a blocking call
3. we are sending value to channel in blocking call. Since the buffer capacity  is 1,
the next value is sent only after 1st value is received in blocking call. 

*/