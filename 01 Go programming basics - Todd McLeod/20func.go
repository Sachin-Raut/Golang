package main

import "fmt"

func main(){
	func (){
		fmt.Println("Anonymous function")
	}()

	func (i int){
		fmt.Println(i)
	}(10)
}