package main

import (
	"fmt"
	"time"
	"math/rand"
)


//lets create 2 channels for 2 different kind of cakes
func main(){
	
	bigCakeChan := make(chan string, 8)
	smallCakeChan := make(chan string, 8)

	go cakeMaker("big cake", 4, bigCakeChan)
	go cakeMaker("small cake", 3, smallCakeChan)

	moreBigCake := true
	moreSmallCake := true

	var cake string

	for moreBigCake || moreSmallCake {

		select {
		case cake, moreBigCake = <- bigCakeChan:
			if moreBigCake {
				fmt.Println("First =", cake)
			}
		case cake, moreSmallCake = <- smallCakeChan:
			if moreSmallCake {
				fmt.Println("Second =", cake)
			}
		}
	}
}

func cakeMaker(kind string, number int, to chan <- string) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < number; i++ {
		//we will sleep upto 500 ms

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		to <- kind
	}
	close(to)
}

/*

1. The output of this program is anywhere between 0 to 500 ms. That is each cake takes
between 0 to 500 ms
2. Let's say we want to execute this process in 250 ms. No matter the no.of cakes 
we generate, we will stop after 250 ms.
3. Add another case

case <- time.After(250 * time.Millisecond):
	fmt.Println("Timed out")
	return
*/