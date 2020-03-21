package main

import "fmt"

func main(){
	s := foo("Mumbai")
	fmt.Println(s)
}

func foo(s string) string {
	a := "Hello " + s
	return a
}