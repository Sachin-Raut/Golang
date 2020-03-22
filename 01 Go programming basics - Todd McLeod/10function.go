package main

import "fmt"

func main(){
	s := foo("Mumbai")
	fmt.Println(s)

	fn, ln := name()
	fmt.Println(fn)
	fmt.Println(ln)

	//variadic means 0 or more parameters
	bar(2,3,4)

	b := []int{4,6,8}
	bar(b...)

	/*
	if you are passing multiple parameters, 
	then variadic parameter should be the last parameter

	func city(a string, b ...int) this is correct
	func city(x ...int, y string) this is wrong
	*/
}

func foo(s string) string {
	a := "Hello " + s
	return a
}

func name()(string, string){
	return "sachin","raut"
}

func bar(x ...int){
	sum := 0
	for _,v := range x {
		sum += v
	}
	fmt.Println(sum)
}