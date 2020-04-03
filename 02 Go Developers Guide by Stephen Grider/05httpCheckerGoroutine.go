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

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}
	
}

func checkLink(link string, ch chan string){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		ch <- "might be down"
		return
	}
	fmt.Println(link, "is up")
	ch <- "link is up"
}