package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	byteSlice, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(byteSlice)

	loginPassword := `password123`

	err = bcrypt.CompareHashAndPassword(byteSlice, []byte(loginPassword))

	if err != nil {
		fmt.Println("Login failed")
		return
	}

	fmt.Println("Logged in")
}