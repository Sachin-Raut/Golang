package main

import (
	"fmt"
	"sync"
)


//communication by sharing memory
func main(){
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printEven(i, &wg)
	}

	wg.Wait()
}

func printEven(x int, wg *sync.WaitGroup){
	if x % 2 == 0 {
		fmt.Printf("%d is even\n", x)
	}
	wg.Done()
}

/*

2. goroutines should communicate via channels
3. channels are like pipes with 2 ends. Messages go in one end when sent, & comes out 
at the other end when requested by the receiver.

*/