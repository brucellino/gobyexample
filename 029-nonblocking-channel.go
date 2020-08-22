// NOTE: https://gobyexample.com/non-blocking-channel-operations

/*
  Base sends and receives on channels are blocking.
  However, we can use select with a default clause to implement nonblocking sends, receives
  and even non-blocking multiway selects.
*/

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool, 0)

	// NOTE: Here is a nonblocking receive.
	// NOTE: If a value is available on messages then select will take that the <- messages
	// NOTE: with that value.
	// NOTE: If not it will immedately take the default case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// NOTE: A nonblocking send works similarly.
	// NOTE: Here msg cannot be sent to the messages channel, because the channel has no
	// NOTE: buffer and there is no receiver.
	// NOTE: Therefore the default case is selected
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// NOTE: We can use multiple cases above the default clause to implement multiway
	// NOTE: nonblocking select. Here we attempt nonblocking receives on both messages
	// NOTE: and signals.
	select {
	case msg := <-messages:
		fmt.Println("receved message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
