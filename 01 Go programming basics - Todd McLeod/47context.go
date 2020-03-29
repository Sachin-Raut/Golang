package main

import (
	"fmt"
	"time"
	"context"
	"runtime"
)

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("error check 1 =", ctx.Err())
	fmt.Println("No of goroutines 1 =", runtime.NumGoroutine())

	go func(){
		n := 0
		for {
			select {
			case <- ctx.Done():
				fmt.Println("context done")
				return
			default:
				n++
				time.Sleep(200*time.Millisecond)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second)

	fmt.Println("error check 2 =", ctx.Err())
	fmt.Println("No of goroutines 2 =", runtime.NumGoroutine())
	fmt.Println("about to cancel context")

	cancel()  // this call executes ctx.Done()

	fmt.Println("cancelled context")

	time.Sleep(time.Second)

	fmt.Println("error check 3 =", ctx.Err())
	fmt.Println("No of goroutines 3 =", runtime.NumGoroutine())
}