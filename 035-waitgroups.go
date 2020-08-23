// NOTE: https://gobyexample.com/waitgroups
/*
To wait for multiple goroutines to finish, we can use a wait group
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	/*
	  This is the function that we'll run in every goroutine.
	  Note that a WaitGroup must be passed to functions by pointer.
	*/

	// NOTE: On return, notify the WaitGroup that we're done
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	// NOTE: Sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// NOTE: This WaitGroup is used to wait for all the goroutines launched here to finish.
	var wg sync.WaitGroup

	// NOTE: Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	// NOTE: Block until the waitGroup counter goes back to 0; all workers notified done.
	wg.Wait()
}

/*
The order of workers starting up and finishing is usually different for each invocation.
*/
