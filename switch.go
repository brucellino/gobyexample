package main

import (
	"fmt"
	"time"
)

// NOTE: https://gobyexample.com/switch
func main() {

	i := 2

	fmt.Println("Write ", i, " as ")
	// NOTE: A basic switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// NOTE:  Use commas to separate multiple expressions in the same case statement
	//        The optional default case is also used here.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// NOTE:   Switch without an expression is another way of using if/else logic
	//         We also how how case expressions can be non-constants -- time.Now()
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// NOTE:  a type switch compares types instead of values
	//        use this to discover the type of an interface value
	//        in this example, variable t will have the type corresponding
	//        to its clause
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("I don't know what type %T is\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
