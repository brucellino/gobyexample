package main

import "fmt"

func main() {
	// Tutorial 6 - basic if/else
	// basic example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Can have if without else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals.
	// Any conditionals declared in the condition are available
	// in the braces
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// no parentheses are needed around statements, but braces are needed in blocks
	// there is no ternary in go.
}
