/*

1. test ensures that your code is doing what you want it to do
2. tests must be in a file that ends with "_test.go"
3. test file should be in the same package as the one being tested.
4. test must be in a function with an signature "func Testxxx(t *testing.T)" 

if you want to test following function 

func mySum(){

}

then test function should be 

func TestMySum(t *testing.T){

}

*/

package main

import "fmt"

func mySum(x ...int) int {
	sum := 0

	for _, value := range x {
		sum += value
	}

	return sum
}

func main() {
	fmt.Println("2+3 =", mySum(2,3))
}