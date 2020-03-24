
/*
Empty inerface

1. interface{}
2. interface that specifies zero methods is known as the "rmpty interface"
3. An empty interface may hold values of any type. They handle value of unknown type.
4. Go interfaces don't enforce a type to implement

*/

package main

import "fmt"

func describe(i interface{}){
	fmt.Printf("value -", i, "%T", i)
}

func main(){
	var i interface{}  //empty interface
	// describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}