/*

There are 2 patterns
1.Fanout
2.Drop

*/

package main

import (
	"fmt"
	"time"
)

func main(){
	emps := 20
	ch := make(chan string, emps) //buffered channel of 20

	for e := 0; e < 20; e++ {
		go func(){
			time.Sleep(100 * time.Millisecond)
			ch <- "paper"
		}()
	}

	for emps > 0 {
		p := <- ch //blocking call
		fmt.Println("emp no", emps, p)
		emps--
	}
}

/*

1. we are creating a buffered channel of 20
2. we are creating goroutines in "for loop". Since for loop iterates 20 times,
20 goroutines will be created
3. "p := <- ch"is a blocking call, unless and until it receives a value, next
statements in the "for loop" won't be executed.

In fanout, we have 20 employees who are sending the 20 values. And we are 
receiving 20 values in another "for loop"

*/