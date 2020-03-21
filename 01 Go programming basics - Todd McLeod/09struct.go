//struct example
package main

import "fmt"


//Person is
type Person struct {
	First string
	Last string
}

//SportsPerson is
type SportsPerson struct {
	Person
	isCricketer bool
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

	sp1 := SportsPerson {
		Person: p1,
		isCricketer: true,
	}

	sp2 := SportsPerson {
		Person : Person {
			First : "Virat",
			Last : "Kohli",
		},
		isCricketer: true,
	}

	fmt.Println(sp2.isCricketer)
	fmt.Println(sp1.isCricketer)

}