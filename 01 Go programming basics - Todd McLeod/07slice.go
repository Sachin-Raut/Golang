package main

import "fmt"

func main(){
	x := []int{1,4,5,2}
	fmt.Println(x)

	for index, value := range x {
		fmt.Println(index,value)
	}
}