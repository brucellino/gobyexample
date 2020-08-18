// NOTE: https://gobyexample.com/closures

package main

import "fmt"

func intSeq() func() int {
	// NOTE: This function returns another function, which we define anonymously in intSeq
	// NOTE: The returned function _closes over_ the variable i to form a closure
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// NOTE: We call intSeq, assigning the result (a function) to nextInt
	// NOTE: This function value captures its own i value, which will be updated each time
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// NOTE: To confirm that the state is unique to that function, create and test a new one
	newInts := intSeq()
	fmt.Println(newInts())
}
