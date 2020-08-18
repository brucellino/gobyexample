package main

import "fmt"

// NOTE: Go supports pointers allowing you to pass references to values and records within
// NOTE: your program

// NOTE: We will demonstrate things with 2 functions: zeroval and zeroptr
func zeroval(ival int) {
	// NOTE: zeroval has an int parameter, so arguments will be passed to it by value
	// NOTE: zeroval will get a copy of ival distinct from the one in the calling function
	ival = 0
}

func zeroptr(iptr *int) {
	// NOTE: zeroptr in contrast has an pointer *int parameter
	// NOTE: The *iptr code in the function body then dereferences the pointer from its
	// NOTE: memory address to the current value at that address
	// NOTE: Assigning a value to a derefenced pointer changes the value at the
	// NOTE: referenced address
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	// NOTE: The &i syntax gives the address in memory of i, ie a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)
	fmt.Println("pointer: ", &i)

}

// NOTE: zeroval doesn't change the i in main, but zeroptr does because it has a reference to the memory address for that variable
