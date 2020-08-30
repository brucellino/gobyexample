// NOTE: https://gobyexample.com/mutexes
/*
  In the previous example, we saw how to manage simple counter state by using atomic operations.
  For more complex state, we can use a mutex to safely access data across multiple goroutines.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// NOTE: For our example, the state will be a map.
	var state = make(map[int]int)
	// NOTE: The mutex will synchronise access to state
	var mutex = &sync.Mutex{}
	// NOTE: We'll keep track of how many read and write opeations we do
	var readOps uint64
	var writeOps uint64

	// NOTE: Here we start 100 goroutines to execute repeated reads against the state
	// NOTE: once per ms in each goroutine.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// NOTE: wait a bit between reads.
				time.Sleep(time.Millisecond)

			}
		}()
	}

	// NOTE: We'll also start 10 goroutines to simulate writes, using the same pattern
	// NOTE: as for reads.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()

				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// NOTE: Let the 10 goroutines work on the state and mutex for a second.
	time.Sleep(time.Second)

	// NOTE: Take and report final operation
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)

	// NOTE: With a final lock of state, show how it ended up
	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}

/*
	Running the program shows that we executed about 19k total operations
	against our mutex-synchronised state.
*/
