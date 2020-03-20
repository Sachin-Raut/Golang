package main

import "fmt"

//for example
func main() {

	/*
	   const a = 10
	   const b = 25.30
	   const city = "Mumbai"
	*/

	/*
	   const (
	   	a = iota // 0
	   	b        // 1
	   	c        // 2
	   )
	*/

	s := "Hello"
	byteString := []byte(s)

	fmt.Println(byteString) // output is ascii values [72 101 108 108 111]

	//for loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//for with single condition
	fmt.Println("for with single condition")
	a := 1
	for a < 5 {
		fmt.Println("a is", a)
		a++
	}

	b := 1

	for {
		if b > 5 {
			break
		}
		fmt.Println("b is", b)
		b++
	}

	c := 1

	for {
		c++
		if c > 100 {
			break
		}
		if c%2 == 0 {
			continue
		}
		fmt.Println("c is", c)
	}
}
