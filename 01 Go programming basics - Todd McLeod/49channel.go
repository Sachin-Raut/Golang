/*

Write a program that launches 10 goroutines & each goroutine adds 10 numbers to a channel.
Then retrieve the numbers from the channel & print them

*/

package main

import "fmt"

func main() {
	ch := make(chan int)

	for j := 0; j < 10; j++ {
		go func(){
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i, <- ch)
	}
	
	fmt.Println("about to exit")
}