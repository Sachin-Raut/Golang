package main

import "fmt"

func main(){
	m := map[string]int {
		"Mumbai":50,
		"Pune":120,
	}

	fmt.Println(m)
	fmt.Println(m["Mumbai"])

	//check if Delhi exists
	v, ok := m["Delhi"]

	fmt.Println(v)

	fmt.Println(ok)

	if _, ok := m["Mumbai"]; ok {
		fmt.Println("Mumbai exists")
	}
}