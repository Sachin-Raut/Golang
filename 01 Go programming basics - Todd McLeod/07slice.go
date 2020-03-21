package main

import "fmt"

func main(){
	x := []int{1,4,5,2}
	fmt.Println(x)

	for index, value := range x {
		fmt.Println(index,value)
	}
	a := []int {1,2,3,4,5,6,7,8,9}

	fmt.Println(a[2:]) //all element from index 2

	fmt.Println(a[1:3]) //all element from index 1 to index 2

	fmt.Println(a[:5]) //all element from index 0 to index 4

	fmt.Println(a[0:4]) //all element from index 0 to index 3

	a = append(a,30,40)

	fmt.Println("a after append 30, 40 is", a)

	b := []int{200,300}

	a = append(a,b...)

	fmt.Println("a after appending b",a)

	y := make([]int, 10, 100)

	fmt.Println("Length of y is ",len(y))
	fmt.Println("Capacity of y is",cap(y))

}