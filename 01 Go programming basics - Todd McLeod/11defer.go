package main

import "fmt"

func main(){
	defer foo()
	bar()
}

func foo(){
	fmt.Println("it's foo")
}

func bar(){
	fmt.Println("it's bar")
}