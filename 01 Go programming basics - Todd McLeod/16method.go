package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	hasLicense bool
}

func (s secretAgent) speak() {
	fmt.Println("I'm", s.first, s.last)
}

func main(){
	sa1 := secretAgent {
		person: person {
			first: "sachin",
			last: "raut",
		},
		hasLicense: true,
	}
	sa1.speak()
}

// func speak() can be called by any variable whose type is "secretAgent"