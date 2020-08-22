// NOTE: https://gobyexample.com/channel-synchronization

/*
  We can use channels to synchronise execution across goroutines.
  Here's an example of using a blocking receive to wait for a goroutine
  to finish.
  When waiting for multiple goroutines to finish, you may before to use
  WaitGroup
*/

package main

import (
	"fmt"
	"time"
)

/*
NOTE: This is a function we'll run in a goroutine.
The done channel will be used to notify another goroutine that this function's
work is done.
*/
func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	// NOTE: Start a worker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)
	// NOTE: Block until we receive a notification from the worker on the channel
	<-done

	// NOTE: If you remove the <- done line from this program, the program would exit
	// NOTE: before the worker ever started.
}
