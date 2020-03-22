/*
create a function that opens 2 files & copies content of 1 file into another
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


/*
This program will work, but there is a bug.
If the call to "os.Create()" fails, then the function will return without
closing the source file. 

Let's fix this bug, using "defer"
*/

func copyFile(fromFile, toFile string){
	src, err := os.Open(fromFile)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(toFile)
	if err != nil {
		return
	}
	defer dst.Close()

	written, err := io.Copy(dst, src)
}

/*
1. "defer" ensures/guarantees that the defer statement would be executed later
2. defer is used to perform clean up actions
3. In this example, defer statement guarantees that the files will be closed 
before returning control out of the function.
4. Let's assume, the call to "os.Create" fails, so it return. But before that
"defer src.Close()" will be executed. 
*/

