// NOTE: https://gobyexample.com/select
/*
  Go's select lets you wait on multiple channel operations.
  Combining goroutines and channels with select is a powerful feature of Go.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// NOTE: We'll select across two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// NOTE: Each channel will receive a value after some amount of time,
	// NOTE: to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(10 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(10 * time.Second)
		c2 <- "two"
	}()

	// NOTE: We'll use select to await both of these values simulataneously,
	// NOTE: printing each one as it arrives
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)

		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

/* NOTE: We receive the values one and then two as expected.
    The total execution time is only ~2 seconds since both sleeps execute concurrently:
time go run 027-select.go
received two
received one
go run 027-select.go  0.20s user 0.03s system 2% cpu 10.177 total
*/
