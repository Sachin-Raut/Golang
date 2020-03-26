package main

import "fmt"

type person struct {
	first string
}

func main(){
	p := person{
		first: "sachin",
	}

	change(&p)
	fmt.Println(p.first) //rahul
}

func change(i *person){
	i.first = "rahul"
}