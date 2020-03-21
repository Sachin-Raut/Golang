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



	/*
	1. Slice is built on top of an array
	2. Slice is dynamic, their size is dynamic
	3. When slice grows, the new array is created & then all the values are copied 
	into the new array & old ones are thrown away. This takes some processing power
	4. If you know how many elements you can store, then you can use "make()" to
	automatically make the underlying array to be big enough size to hold all the 
	elements you want to store. This saves processing power of the compiler.
	*/

	y := make([]int, 10, 100)

	fmt.Println("Length of y is ",len(y))
	fmt.Println("Capacity of y is",cap(y))

}