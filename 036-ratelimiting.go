// NOTE: https://gobyexample.com/rate-limiting
/*
Rate limiting is an important mechanism for controlling resource utilisation and
maintaining quality of service.
Go elegantly supports rate limiting with goroutines, channnels and tickers
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// NOTE: First, we'll look at basic rate limiting.
	// NOTE: Suppose we want to limit ouir handling of incoming requests.
	// NOTE: We'll serve these requests off a channel of the same name

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// NOTE: This limiter channel will receive a value every 200 ms.
	// NOTE: This is the regulator rate limiter scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// NOTE: By blocking on a recive from the limiter channel before serving each request
	// NOTE: we limit ourselves to 1 request every 200ms

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// NOTE: We may want to allow short bursts of requests in our rate limiting scheme
	// NOTE: while preserving the overall rate limit.
	// NOTE: We can accomplish this by buffering our limiter channel.
	// NOTE: The burstyLimiter channel will allow bursts of up to 3 events.

	burstyLimiter := make(chan time.Time, 3)
	// NOTE: Filling up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// NOTE: Every 200 ms we will try to add a new value to burstyLimiter, up to it's limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// NOTE: Now simluate 5 more incoming requests.
	// NOTE: The first 3 of these will benefit from the capability of burstyLimiter
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Request ", req, time.Now())
	}
}

/*

Running our program, we see the first batch of requests handled every 200ms as  desired

go run 036-ratelimiting.go
request 1 2020-08-23 12:46:19.993109032 +0200 CEST m=+0.200365658
request 2 2020-08-23 12:46:20.193105436 +0200 CEST m=+0.400362165
request 3 2020-08-23 12:46:20.393124752 +0200 CEST m=+0.600381410
request 4 2020-08-23 12:46:20.593092908 +0200 CEST m=+0.800349532
request 5 2020-08-23 12:46:20.793009462 +0200 CEST m=+1.000266034

for the second batch, we serve the first 3 immediately, then rate limit
Request  0 2020-08-23 12:46:20.793103797 +0200 CEST m=+1.000360271
Request  1 2020-08-23 12:46:20.793113867 +0200 CEST m=+1.000370342
Request  2 2020-08-23 12:46:20.793119249 +0200 CEST m=+1.000375722
Request  3 2020-08-23 12:46:20.993235549 +0200 CEST m=+1.200492028
Request  4 2020-08-23 12:46:21.193309554 +0200 CEST m=+1.400566032

*/
