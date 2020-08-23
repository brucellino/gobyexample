// NOTE: https://gobyexample.com/range-over-channels

/*
In a previous example we saw how for and range provide iteration over basic data structures.
We can also use this syntax to iterate over values received from a channel.
*/

package main

import "fmt"

func main() {
	// NOTE: We'll iterate over two values in the channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// NOTE: This range iterates over each element as it's received from the queue.
	// NOTE: because we closed the channel above, the iteration terminates after receiving
	// NOTE: two elements
	for elem := range queue {
		fmt.Println(elem)
	}
}

/*
This example also showed that it's possible to close a non-empty channel but still have
the remaining values be received.
*/
