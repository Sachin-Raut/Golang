//basic benchmark test

package basictest

import (
	"fmt"
	"testing"
)

var gs string

//BenchmarkSprint() tests the performance of using Sprint
func BenchmarkSprint(b *testing.B){
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("Hello")
	}
	gs = s
}


//BenchmarkSprintf() tests the performance of using Sprint
func BenchmarkSprintf(b *testing.B){
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("Hello")
	}
	gs = s
}



/*
Benchmarking in Go is extremely powerful

1. here we are benchmarking CPU profile of how fast these "Sprint()" & "Sprintf()" runs
2. benchmarking makes sure that we are not guessing
3. by default, the benchmark test runs for 1 second. But we are going to increase it to 3 seconds
Just to make sure we get enough iterations. We are increasing the benchtime to get more iterations 
to feel more confident

4. $ go test -run none -bench . -benchtime 3s
(right now we are looking at only CPU)

5. $ go test -run none -bench . -benchtime 3s -benchmem
(right now we are looking at CPU & memory allocation)

(we are using "none" because there are no test functions,
-bench . for all the benchmarks)

output -

BenchmarkSprint-2       21594752               153 ns/op               5 B/op          1 allocs/op
BenchmarkSprintf-2      29274661               116 ns/op               5 B/op          1 allocs/op
PASS

(1st is 153 ns/op & 5 bytes allocated over 1 object)
(2nd is 116 ns/op & 5 bytes allocated over 1 object) 

Conclusion = 

Sprintf() is faster than Sprint()

*/

/*

1. you must always validate benchmarks to know that the results are accurate
(we can't blindly trust the result, because our machine might not be idle)

2. run each benchmark separately & then check the results. This gives accurate results.

*/