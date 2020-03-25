//returning a function

package main

import "fmt"

func main()  {
	x := func() int {
		return 5
	}()

	fmt.Println(x)      // 5

	y := bar()

	i := y()

	fmt.Println(i)     // 10
}

func bar() func() int {
	return func() int {
		return 10
	}
}