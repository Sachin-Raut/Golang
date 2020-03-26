package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int {5,2,9,4,6}
	sort.Ints(s)
	fmt.Println(s)
}