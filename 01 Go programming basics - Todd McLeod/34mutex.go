/*

1. in previous example we had created race condition, 
2. we will solve race condition using mutex
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

	var mu sync.Mutex

	for i := 0; i < gs; i++ {

		go func(){
			mu.Lock()

			v := counter
			time.Sleep(time.Millisecond)
			v++
			counter = v

			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines -",runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Last goroutine count-", runtime.NumGoroutine())
	fmt.Println("count -",counter)
}