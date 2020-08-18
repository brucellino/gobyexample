package main

import (
	"fmt"
	"math"
)

// NOTE: https://gobyexample.com/interfaces

// NOTE: here is a basic interface for geometric shapes

type geometry interface {
	area() float64
	perim() float64
}

// NOTE: for our example, we'll implement this interface on rectand circle types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// NOTE: To implement an interface in go, we need to implement the methods in the interface

// NOTE: Here, we impement geometry on rects:
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// NOTE: Now, implement geometry for circle
func (c circle) area() float64 {
	return 2 * math.Pi * (c.radius * c.radius)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// NOTE: If a variable has an interface type, then we can call methods that are in the named interfaces
// NOTE: Here is a generic measure function taking advantage of this to work on any geometry

func measure(g geometry) {
	fmt.Println("Measuring ", g)
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
