/*
create a function that opens 2 files & copies content of 1 file into another
*/


/*
This program will work, but there is a bug.
If the call to "os.Create()" fails, then the function will return without
closing the source file. Let's fix this problem in next program.
*/


package main

import (
	"fmt"
	"os"
)
func main(){
	copyFile(fromFile, toFile)
}

func copyFile(fromFile, toFile string) {
	src, err := os.Open(fromFile)
	if err != nil {
		return
	}

	dst, err := os.Create(toFile)
	if err != nil {
		return
	}

	written, err := io.Copy(dst,src)

	dst.Close()
	src.Close()
}