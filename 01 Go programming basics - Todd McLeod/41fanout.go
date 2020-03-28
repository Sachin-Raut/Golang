package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go populate(ch1)

	go fanout(ch1, ch2)

	for v := range ch2 {
		fmt.Println(v)
	} 
	fmt.Println("about to exit")
}

func populate(ch chan int){
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func fanout(ch1, ch2 chan int) {
	var wg sync.WaitGroup

	for v := range ch1 {
		wg.Add(1)

		go func(v2 int){
			time.Sleep(time.Millisecond * 50)
			ch2 <- 20
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(ch2)
}