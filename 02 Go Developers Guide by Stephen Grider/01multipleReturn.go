package main

import "fmt"

//example of multiple return values

func main() {

	a, b := multipleReturn()
	fmt.Println("a =", a, "\nb =", b)

}

func multipleReturn() (int, int) {
	return 2,6
}
