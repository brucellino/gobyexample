// NOTE: https://gobyexample.com/timers

/*
  We often want to execute go code at some point in the future, or repeatedly at some interval.

  Go's built-in timer and ticker features make both of these tasks easy.
  We'll look first at timers and then at tickers.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// NOTE: Timers represent a single event in the future.
	// NOTE: You tell the timer how long you want to wait and it provides a channel
	// NOTE: that will be notified at that time.
	// NOTE: This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// NOTE: The <- timer1.C blocks on teh timer's chanel C until it sends a value indicating
	// NOTE: that the timer has fired.
	<-timer1.C
	fmt.Println("Timer 1 has fired.")

	// NOTE: If you just wanted to wait, you could have used time.Sleep.
	// NOTE: One reason a timer may be useful is that you can cancel the timersnote
	// NOTE: before it fires. eg:

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 has fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 has  stoppped.")
	}
	// NOTE: Give the timer 2 enough time to fire, if it was ever going to, to show that
	// NOTE: it is in fact stopped.
	time.Sleep(2 * time.Second)
}

/*
  The first timer will fire ~2s after we start the program, but the second should be
  stopped before it has a chance to fire.
*/
