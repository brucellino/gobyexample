package main

import "fmt"

func main() {

	// Basic loop with single condition and starting point.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic initial / condition / after loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop  until it reaches break or
	// a return is received from the enclosing function
	for {
		fmt.Println("loop")
		break
	}

	// Can also continue to the next iteration of the loop based on a condition
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
