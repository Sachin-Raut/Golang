package main

import "fmt"

func main(){
	a := 20

	fmt.Println(a)    //20
	fmt.Println(&a)   // memory address of a

	b := &a

	fmt.Println(*b)   //20

	*b = 45

	fmt.Println(a)    //45
}