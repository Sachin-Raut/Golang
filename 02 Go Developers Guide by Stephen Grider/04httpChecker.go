package main

import (
	"fmt"
	"net/http"
)

func main(){
	links := []string {
		"http://google.com",
		"http://facebook56.com",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	fmt.Println(link, "is up")
}

/*

1.currently our program is executing in sequential order & hence there's a delay 
in executing the entire program.
2.this is wrong approach. Let's use goroutines to execute multiple links simultaneously
3.concurrency isn't parallelism
4.concurrency doesn't guarantee parallelism

*/