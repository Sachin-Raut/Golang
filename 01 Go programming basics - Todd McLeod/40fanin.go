/*

1. In fanin pattern, we will take values from 2 channels & put them on 1 channel

*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

//send channel
func send(even, odd chan <- int){

	for i := 0; i < 20; i++ {
		if i % 2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

//receive channel (even & odd are receive channels, fanin is send channel)
func receive(even, odd <- chan int, fanin chan <- int){
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func(){
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}