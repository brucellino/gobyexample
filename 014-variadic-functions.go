// NOTE: https://gobyexample.com/variadic-functions

package main

// NOTE: Variadic functions can be called with any number of trailing functions

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	// NOTE: If you already have multiple arguments in a slice, apply them for a variadic
	// NOTE: function using func(slice ...) like this
	sum(nums...)
}
