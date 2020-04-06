/*

1. In Go, we don't start OS thread for each goroutine
2. We will have for eg. 10 threads, operating 200 goroutines
3. Goroutines are generally much greater than thread
4. Each goroutine uses 2 kb memory, whereas OS thread uses 2 MB memory.
5. So 1000 goroutines, utilizes the same memory as 1 OS thread.

*/