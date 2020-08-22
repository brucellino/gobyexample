// NOTE: https://gobyexample.com/channels

/*
  Channels are the pipes that connect concurrent goroutines.
  You can send values into chanbnels from one goroutine and
  receive those values into another goroutine.
*/

package main

import "fmt"

func main() {

	// NOTE: Create  anew cchannel with make(chan val type)
	// NOTE: Channels are typed by the values they convey.
	messages := make(chan string)

	// NOTE: Send a value into a channel using the
	// NOTE: channel <- syntax.
	// NOTE: Here we send a "ping" to the "messages" channel we created above,
	// NOTE: from a new goroutine.
	go func() {
		messages <- "ping"
	}()

	// NOTE: The <- channel syntax receives a vlaue from the channel.
	// NOTE: Here we'll receive the ping mesasge we sent above and pring it out.
	msg := <-messages
	fmt.Println(msg)

}

/*
  When we run the program, the "ping" message is successfully passed from one goroutine to another via our channel.

  By default, sends and receives block until both the sender and receiver are ready.
  This property allowed us to wait at the end of our program for the ping message without
  having to use any other synchronisation.
*/
