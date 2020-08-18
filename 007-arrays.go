// Note: https://gobyexample.com/arrays

package main

import "fmt"

func main() {

	// NOTE:  Here we create an array that holds exactly 5 ints
	//        The types of elements and length are both part of the array's types
	//        By default, the array is zero-valued.
	//        For ints, this is 0
	var a [5]int
	fmt.Println("emp:", a)

	// NOTE:  We can set a value at an index using array[index] = value
	//        and get a value using array[index]
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	// NOTE:  the builtin len function returns the length of an array
	fmt.Println("len:", len(a))

	// NOTE:  Declaring an array with length, type and content
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// NOTE:  Array types are 1-D, but you can compose types
	//        to build multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
