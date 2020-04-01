package main

import "fmt"
import "strings"

//example of multiple return values

func main() {

	a, b := multipleReturn()
	fmt.Println("a =", a, "\nb =", b)


	cityName := "Mumbai"
	fmt.Println([]byte(cityName)) //array of ascii values byteArray

	//Joining the slice of strings
	 slice1 := []string{"Mumbai","Pune"}
	 slice2 := strings.Join(slice1,",")
	 fmt.Println(slice2)

}

func multipleReturn() (int, int) {
	return 2,6
}
