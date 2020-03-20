package main

import "fmt"

//zero values
var a int       //by default a is 0
var city string // by default city is ""

//create your own type
type hotdog int //hotdog underlying type is int
var b hotdog

func main() {
	//short declaration operator
	x := 42

	fmt.Println(x)
	b = 50

	/*
		a = b is not possible because Golang is static language
		& hence type casting is not possible. We have to do conversion
	*/

	a = int(b)

	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
