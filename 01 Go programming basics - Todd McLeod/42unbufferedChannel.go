package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)
	a := make(chan int)

	go func(){
		p := <- ch //(receive operation) this is blocking call

		//now perform another tasks
		time.Sleep(time.Second)
		fmt.Println(p)
		fmt.Println("hi")

		a <- 15

	}()

	time.Sleep(2 * time.Second)

	ch <- 20 //(send operation)
	fmt.Println(<-a)

}

/*
1. we are creating channel (unbuffered)
2. we are creating goroutine. This goroutine has receive channel operation. This is blocking call.
Unless and until this receives a value, Go scheduler doesn't execute next tasks in goroutine.
3. in main goroutine, we are delaying by 2 seconds(i.e latency is 2 seconds)
4. from main goroutine we are sending value 20 to the channel.

After both the "send channel & receive channel" operations are executed, the Go scheduler can choose
to execute any statement it wants (it may choose to execute task from main goroutine or it may choose
to execute task from another goroutine)

This means here using "Print" statements can fool us about order of execution. 
In such scenarios don't be dependent on "print" statement

In unbuffered channel, there is a guarantee that the value being sent was received, 
but the delay or latency is unknown.

*/