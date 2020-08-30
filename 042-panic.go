// NOTE: https://gobyexample.com/panic

/*
  A panic typically means something when unexpectedly wrong.
  Mostly, we use it to fail fast on errors that shouldn't ocur during normal operations,
  or that we aren't prepared to handle gracefully.
*/

package main

import "os"

func main() {
	panic("a problem")

	/*
	   A common use of panic is to abort if a function returns an error value that we don't
	   know how to handle.
	   Here we panic if we get an unexpected error creating a new file:
	*/
	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}
}

/*
  Unlike some languages which use exceptions for handling errors, in go, it's idiomatic
  to use error-indicating return values wherever possible.
*/
