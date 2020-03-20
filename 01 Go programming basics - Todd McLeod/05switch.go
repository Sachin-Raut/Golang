package main

import "fmt"

func main() {
	switch "Mumbai" {
	case "Pune", "Delhi":
		fmt.Println("It's Pune or Delhi")
	case "Mumbai":
		fmt.Println("It's Mumbai")
	default:
		fmt.Println("It's default")
	}
}
