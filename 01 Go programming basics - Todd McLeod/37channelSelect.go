// we will see how to range over a channel


/*
send-only channel
ch := make(chan <- int, 2)

receive-only channel
ch := make(<- chan int, 2)

general channel
ch := make(chan int)
*/

package main

import "fmt"

func main(){
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	quitChannel := make(chan int)

	//send channel
	go send(evenChannel, oddChannel, quitChannel)

	//receive channel
	receive(evenChannel, oddChannel, quitChannel)
}

func send(e, o, q chan <- int) {
	for i := 0; i < 100; i++ {
		
		//put even number on evenChannel
		if i%2 == 0 {
			e <- i
		} else {
		
		//put odd number on oddChannel
		o <- i
		}		
	}
	q <- 0
}

func receive(e, o, q <- chan int){
	for {
		select {
		case v := <- e:
			fmt.Println("even channel -", v)
		case v := <- o:
			fmt.Println("odd channel -", v)
		case v := <- q:
			fmt.Println("quit channel -", v)
			return
		}
	}
}