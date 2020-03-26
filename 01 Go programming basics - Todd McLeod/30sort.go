package main

import (
	"fmt"
	"sort"
)

//Person is
type Person struct {
	first string
	age int
}

//ByAge is
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i,j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i,j int) bool {
	return a[i].age < a[j].age
}

func main(){
	p1 := Person{"sachin1", 20}
	p2 := Person{"sachin2", 25}
	p3 := Person{"sachin3", 10}
	p4 := Person{"sachin4", 28}

	people := []Person{p1,p2,p3,p4}

	fmt.Println(people)

	sort.Sort(ByAge(people))

	fmt.Println(people)

}