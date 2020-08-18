package main

// NOTE: https://gobyexample.com/errors

/*
  In go, it's idiomatic to communicate errors via an explicit, separate return value
  This constrasts with the exceptions used in languages like Java and Ruby and the
  overloaded single result/error value somtimes used in C.
  Go's approach makes it easy to see which functions return errors and to handle them
  using the same language constructs used for other non-error tasks
*/

import (
	"errors"
	"fmt"
)

// NOTE: By convention, errors are the last return value
// NOTE: and have type error, a builtin interace
func f1(arg int) (int, error) {
	if arg == 42 {
		// NOTE: errors.New constructs a basic error value with the error message
		return -1, errors.New("can't work with 42")
	}

	// NOTE: a nil alue in the error position indicates that there was no error
	return arg + 3, nil
}

/*
it's possible to use custom types as rrors, by implementing Error() method on them.
Here is a variant on the eample above tht uses a custom type to explicitly repreent
an argument error
*/
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// NOTE: In this case we use &argError syntax to build a new struct.
		// NOTE: We supply values for the two fields arg and prob
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 ailed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
