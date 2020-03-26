package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
}

func main() {

	jsonString := `[{"First":"Sachin"},{"First":"Leonardo"}]`

	byteSlice := []byte(jsonString)

	fmt.Printf("%T\n",jsonString)

	fmt.Printf("%T\n",byteSlice)

	var people []person

	err := json.Unmarshal(byteSlice, &people)

	if err != nil {
		fmt.Println(err)
	}

	for index, value := range people {
		fmt.Println("Person number",index)
		fmt.Println(value.First)
	}
}