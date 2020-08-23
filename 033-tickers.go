// https://gobyexample.com/tickers
/*
  Timers are for when you want to do something once in the future.
  Tickers are for repeatedly doing something at regular intervals.
  Here's an example of a ticker that ticks periodically until we stop it.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	  Tickers use a similar mechanism to timers:
	  a channel that is sent values.
	  Here, we'll use the select builtin on the channel to await the values as they arrive
	  every half second.
	*/

	ticker := time.NewTimer(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()

	// NOTE: Tickers can be stopped like timers.
	// NOTE: Once a ticker is stopped, it won't receive any more values on it's channel.
	// NOTE: We'll stop ours after 1600ms.

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

/*
  When we run this program the ticker shouold tick 3 times before we stop it.
*/
