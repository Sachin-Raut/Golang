package main

import "fmt"

//panic, defer & recovery

func main() {
	f()
	fmt.Println("returned normally from f()")
}

func f() {
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()
	fmt.Println("calling g")

	g(0) //here "panic" stops the execution of f(), hence next statements are not executed

	fmt.Println("returned normally from g")
}

func g(i int){
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i)) // i is 4
	}

	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i+1)
}

/*

calling g
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
recovered in f 4
returned normally from f()

*/

/*

Explanation

1. "panic" stops the execution of f()
2. but before that it executes all the defer statements.
3. after that it returns to its caller (in this case "main")
4. "recover" is only useful inside "deferred function"
5. during normal execution, a call to "recover" returns nil & have no other effect
6. if the current goroutine is panicking, a call to recover will capture 
   the value given to panic & resume normal execution.
7. if we remove "defer func" from f(), then the panic is not recovered 
   & reaches top of the goroutines call stack & terminates the program.
8. "defer" is also used for releasing a "mutex"

mu.Lock()
defer mu.Unlock()

*/