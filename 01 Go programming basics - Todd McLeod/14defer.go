package main

import "fmt"

func main(){

	//deferred function calls are executed in last in first out order (LIFO)
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("Hi")

	aa := a()
	fmt.Println(aa)
}

func a() (i int) {
	defer func(){
		i++         //output is 11
	}()
	return 10
}

/*

Output is

Hi 
11
defer 2
defer 1


*/