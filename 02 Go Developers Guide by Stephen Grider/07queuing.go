package main

import (
	"fmt"
	"sync"
	"net/http"
)

//queuing
func main(){

	var wg sync.WaitGroup
	work := make(chan string, 100)

	numWorker := 100 // 100 goroutines

	for i := 0; i < numWorker; i++ {
		go webGetworker(work, &wg)
	}

	//6 array elements

	urls := [6]string{
		"http://facebook.com",
		"http://apple.com",
		"http://instagram.com",
		"http://flipkart.com",
		"http://amazon.com",
		"http://twitter.com",
	}

	//lets fetch each url 50 times, i.e total 300 fetches

	for i := 0; i < 50; i++ {
		for _, url := range urls{
			wg.Add(1)
			work <- url
		}
	}
	wg.Wait()
}

func webGetworker(in <- chan string, wg *sync.WaitGroup) {
	for {
		url := <- in

		res, err := http.Get(url)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Get", url, "Status", res.StatusCode)
		}
		wg.Done()
	}
}

/*

1. we are fetching 300 urls, on 100 goroutines
2. if we fetch 300 urls on 300 goroutines, then the result will be faster
3. if u increase goroutines to 600, the result speed will still be pretty much same. 
But extra 300 goroutines will add cost to efficiency & this should be avoided.

*/