package main

import "fmt"

func main() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	quitChannel := make(chan bool)

	go send(evenChannel, oddChannel, quitChannel)

	receive(evenChannel, oddChannel, quitChannel)

	fmt.Println("about to exit")
}

//send channel.
func send(even, odd chan <- int, quit chan <- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

//receive channel.
func receive(even, odd <- chan int, quit <- chan bool) {
	for {
		select {
		case v := <- even: 
			fmt.Println("even-",v)
		case v := <- odd:
			fmt.Println("odd-",v)
		case i, ok := <- quit:
			if ok {
				fmt.Println("from comma ok", i, ok)				
			} else {
				fmt.Println("from comma okkkk", i)				
			}
			return
		}
	}
}