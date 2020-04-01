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

5. In terminal, execute "$ go test" or "$ go test -v"

*/

package main

import "testing"

func TestMySum(t *testing.T){
	
	/*

	//this one is for testing single test data
	x := mySum(2,3)
	if x != 5 {
		t.Error("Expected 5, got",x)
	}

	*/


	//this one is for testing multiple test data
	type test struct {
		data []int
		answer int
	}

	tests := []test {		
		test{[]int{10,11},21},
		test{[]int{10,10},20},
	}

	for _,v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected ",v.answer, "Got ",x)
		}
	}
}