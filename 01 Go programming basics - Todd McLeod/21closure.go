package main

import "fmt"

func main() {
	{
		//y can only be acessed within inner code block
		y := 5
		fmt.Println(y)
	}
}