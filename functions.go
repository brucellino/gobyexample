// NOTE: https://gobyexample.com/functions

package main

import "fmt"

// NOTE: functions are central in go.

func plus(a int, b int) int {
	// NOTE: the function plus takes two ints and returns an int
	// NOTE: go requires explicit returns - unlike Ruby
	return a + b
}

func plusPlus(a, b, c int) int {
	// NOTE: When you have multiple consecutive parameters of the same type, you can omit
	// NOTE: the type name for the like-typed parameters up to the final parameter that
	// NOTE: declares the type
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
