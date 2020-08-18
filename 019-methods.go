package main

import "fmt"

// NOTE: https://gobyexample.com/methods
// NOTE: Go supports methods on struct types

type rect struct {
	width, height int
}

func (r *rect) area() int {
	// NOTE: This area method has a receiver type of *rect
	return r.width * r.height
}

func (r rect) perim() int {
	// NOTE: Methods can be defined for either pointer or value receiver of our struct
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// NOTE: We call the 2 methods defined for our structs
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// NOTE: Go automatically handles conversion between values and pointers for method calls
	// NOTE: you may want to use a pointer receiver type to avoid copying on method calls
	// NOTE: or allow the method to mutate the receiving struct
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
