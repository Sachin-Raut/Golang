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

runtime.GOMAXPROCS(2)

7. WaitGroup is used to manage concurrency

5. Benchmark

$ go test -run none -bench . -benchtime 3s

(we are running all the benchmarks & we are increasing benchtime to 3 seconds from 1 second)
(we are not running any tests, hence "none")

*/