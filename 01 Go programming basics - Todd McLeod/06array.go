package main

import "fmt"

//array example
func main(){
	var x [5]int
	x[2] = 51
	fmt.Println(x)
	fmt.Println(len(x))
}