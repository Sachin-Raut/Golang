/*

1. we will cerate race condition
2. we will launch multiple goroutines & all those goroutines will access same variable
3. "go run -race main.go" (for checking race condition)

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	counter := 0

	const gs = 4

	var wg sync.WaitGroup

	wg.Add(gs)	//we are creating 4 goroutines

	for i := 0; i < gs; i++ {
		go func(){
			v := counter
			time.Sleep(time.Second)
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines -",runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Last goroutine count-", runtime.NumGoroutine())
	fmt.Println("count -",counter)
}