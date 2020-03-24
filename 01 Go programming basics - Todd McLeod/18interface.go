/*

1. interface allows to do polymorphism
2. interface allows to define behaviour

*/

package main

import "fmt"

type person struct {
	first string
	last string
}

type human interface {
	speak()
}

func (p person) speak(){
	fmt.Println("I'm from speak()")
}

func foo(h human) {
	fmt.Println("I'm from human")
}

/*
1.any variable of type "person" is also of type "human"
because person is implementing speak
*/

func main(){

	p := person {
		first:"sachin",
		last:"raut",
	}

	foo(p)

	p.speak()
}