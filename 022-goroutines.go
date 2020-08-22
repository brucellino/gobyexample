package main

// https://gobyexample.com/goroutines

// NOTE: a goroutine is a lightweight thread of execution
import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// NOTE: Suppose we have a functionl call f(s).
	// NOTE: Here's how we'd call that in the usual way, running it synchronously:
	f("direct")

	// NOTE: To invoke this function in a goroutein, use go f(s).
	// NOTE: This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	// NOTE: You can also start a goroutine for an anonymous function call:
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// NOTE: Our two function calls are runnning asynchronously in separate goroutines now.
	// NOTE: Wait for them to finish.
	time.Sleep(time.Second)

	fmt.Println("done")

	// NOTE: When we run this program , we see the output of the blockjing call first,
	// NOTE: then the output of the two goroutines..
	// NOTE: This interleaving reflects the goroutines being run concurrently by the
	// NOTE: go runtime.
}
