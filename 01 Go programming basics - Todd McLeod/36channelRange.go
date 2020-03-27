/*

1. range keeps retrieving from channel until the channel is "close"
2. if we don't "close", then range waits for more values to come

*/

package main

import "fmt"

func main() {
	ch := make(chan int)

	//send channel (sending values to the channel)
	go func(){
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	//receive
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}