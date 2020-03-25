package main

import "fmt"

func main(){
	a := factorial(5)
	fmt.Println(a)

	b := loopFact(5)
	fmt.Println(b)
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}

	return i * factorial(i - 1)
}

func loopFact(i int) int {
	total := 1

	for ; i > 0; i-- {
		total *= i 
	}
	return total
}