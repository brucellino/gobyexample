// NOTE: Go has builtin support for multiple return values.
// NOTE: This feature is often used in idiomatic go, for example to return
// NOTE: both a result and an error value from a function

package main

import "fmt"

// NOTE: The (int, int) in this function signature shows that the function
// NOTE: returns two ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	// NOTE: If you're only interested in one of the return variables, ignore the other with _
	_, c := vals()
	fmt.Println(c)
}
