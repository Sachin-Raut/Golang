package main

import (
	"fmt"
	"encoding/json"
) 

type person struct {
	First string
}

func main()  {
	p1 := person {
		First: "Sachin",
	}

	p2 := person {
		First: "Leonardo",
	}

	people := []person{p1,p2}

	fmt.Println(people)

	//if we want to send "people", then we need to marshal it to convert to JSON

	byteSlice, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteSlice))
}