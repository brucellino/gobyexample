// NOTE: https://gobyexample.com/slices
// NOTE:  slices are a key data structure in Go, giving a more powerful interfce than arrays
package main

import "fmt"

func main() {
	// NOTE:  Unlike arrays, slices are typed only by the elements they contain
	//        not by the number of elements
	//        To createan empty slice with nonzero length, use the builtin make function
	//        Here we make a slice of strings of length 3 (initially zero-valued)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// NOTE: We can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	// NOTE: len returns the length of the slice as expected
	fmt.Println("len:", len(s))

	// NOTE: In addition to these basic operations, slices also support several more
	//        - the builtin append - which returns a slice with one or more new values
	// NOTE: We nee to accept a return value from the append function as we may get a new slice

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// NOTE: Slices can also be copy'd.
	// NOTE: here we create an empty slice c of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// NOTE: Slices support a slice operator with the syntax slice[low:high] which selects a
	// NOTE: slice of a slice.
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	// NOTE: We can declare and initialise a variable fora slice in a single line
	t := make([][]int, 3)
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		// NOTE: slices can be composed into multidimensional data structures.
		// NOTE: The length of the inner slices can vary, unlike with multi-D arrays
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
