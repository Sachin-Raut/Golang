//struct example
package main

import "fmt"


//Person is
type Person struct {
	First string
	Last string
}

func main(){
	p1 := Person{
		First:"Sachin",
		Last:"Tendulkar",
	}

	p2 := Person{
		First:"Rahul",
		Last:"Dravid",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.First)
	fmt.Println(p2.First)
}