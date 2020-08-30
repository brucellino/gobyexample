// NOTE: https://gobyexample.com/atomic-counters

/*
The primary mechanism for managing state in Go is communication over channels.
We saw this for example in worker pools.
There are a few other options for managing state though.
Here we'll look at using the sync/atomic pacakge for atomic counters accessed by
multiple goroutines.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// NOTE: We'll use an unsigned integer to represent our (>0) counter
	var ops uint64
	// NOTE: A waitGroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup
	count := 1000
	// NOTE: We'll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// NOTE: To atomically increment the counter we use AddUint64,
		// NOTE: Giving it the memory address of the ops counter with the & syntax
		go func() {
			for c := 0; c < count; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Ops: ", ops)
}

/*
  We expect to get exactly 50k operations.
  Had we used the non-atomic ops + to increment the counter, we'd likely get a different
  number, changing between runs, because the goroutines would interfere with each other.
  Moreover, we'd get dat race failures when running with the -race flag.

*/
