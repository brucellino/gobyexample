// NOTE: https://gobyexample.com/channel-buffering

/*
  NOTE: By default channels are unbuffered, meaning that they will only accept
  sends ( chan <- ) if there is a corresponding receive(<- chan) ready to
  receive the sent value.

  Buffered channels accept a limited number of values without a corresponding
  receiver for those values.

*/

package main

import "fmt"

func main() {
	// NOTE: Here we make a channel of strings buffering up to two values
	messages := make(chan string, 2)

	// NOTE: Because this channel is buffered. we can sebd these values into the channel
	// NOTE: without a concurrent receive

	messages <- "buffered"
	messages <- "channel"

	// NOTE: later, we can receive these two values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
