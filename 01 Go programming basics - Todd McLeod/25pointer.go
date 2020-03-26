package main

import "fmt"

func main(){
	a := 20
	fmt.Println(a) // 20
	foo(&a)
	fmt.Println(a) // 30
}

func foo(y *int){
	fmt.Println("y before change",*y) //20
	*y = 30
}