/*

1. Type is everything in Golang. It's life.
2. Go is a pass by value language.
3. It doesn't have constructor
4. Go source code is UTF-8 (Unicode transformation format)

5. Caching

Reading from L1 = super-duper fast (top speed) 
Reading from L2 = relatively fast (second top)
Reading from L3 = normal (third top)
Reading from main memory is slow (fourth top)

6. Let's say you have 4 cores and want to run code on 2 cores, then do as follows.

func init(){

	runtime.GOMAXPROCS(2)

}

7. WaitGroup is used to manage concurrency 

wg.Wait() - provides guarantee about synchronisation
(wait() orders go scheduler to wait & scheduler has to obey the orders.)

runtime.Gosched() - doesn't provide guarantee about synchronisation 
(gosched() requests go scheduler to wait, & scheduler may or may not accept the request. 
Hence gosched isn't guaranteedThis should mostly be used while running tests)

8. Benchmark

$ go test -run none -bench . -benchtime 3s

(we are running all the benchmarks & we are increasing benchtime to 3 seconds from 1 second)
(we are not running any tests, hence "none")

*/