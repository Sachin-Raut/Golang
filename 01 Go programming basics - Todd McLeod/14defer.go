package main

import "fmt"

func main(){

	//deferred function calls are executed in last in first out order (LIFO)
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("Hi")
}

/*

Output is

Hi 
defer 2
defer 1

*/