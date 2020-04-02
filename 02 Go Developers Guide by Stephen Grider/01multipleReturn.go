package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

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

	 //save data to file "city_name" on hard drive
	 ioutil.WriteFile("01multiple", []byte("abcd"), 0666)

	 //read data from file "city_name" on hard drive
	 byteSlice, err := ioutil.ReadFile("01multiple")
	if err != nil {
		fmt.Println("Error -", err)
		os.Exit(1)

		/*
		Exit(0) - it indicates success
		Exit(1) - it indicates an error
		
		Exit terminates the program immediately, "deferred functions" are not executed
		*/
	}
	fmt.Println("byte slice =",byteSlice)
	normalString := string(byteSlice)
	fmt.Println(normalString)

}

func multipleReturn() (int, int) {
	return 2,6
}
