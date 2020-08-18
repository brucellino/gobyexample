package main

import "fmt"

// NOTE: Classic factorial example
func fact(n int) int {
	// NOTE: This function calls itsel unless it reaches the base case
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
