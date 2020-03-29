/*

1.context is a tool that can be used with concurrent design patterns to make sure that
if you have a process/task, that launches other goroutines. When you cancel the proces/task, 
then all of the launched goroutines are also cancelled.

2. we do this so that we don't leak goroutines

3. leaking goroutines uses resources. let's say we have a process & as part of that process
we launch no.of goroutines. And when we close the process, the goroutines are still running 
& using the resources.

4. context can be used to address the problem of leaking goroutines issue.

5. context is used to cancel a running goroutine.

*/

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context - \t", ctx)
	fmt.Println("context error - \t", ctx.Err())
	fmt.Printf("context type - \t%T\n", ctx)


	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context - \t", ctx)
	fmt.Println("context error - \t", ctx.Err())
	fmt.Printf("context type - \t%T\n", ctx)

	cancel()

	fmt.Println("cancel - \t", cancel)
	fmt.Println("context error - \t", ctx.Err())
	fmt.Printf("context type - \t%T\n", ctx)

}