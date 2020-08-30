// NOTE: https://gobyexample.com/sorting

/*
  Go's sort package implements sorting for builtins and user-defined types.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
	  Sort methods are specific to the builtin type;
	  here's an example for strings.
	  Note that sorting is in-place, so it changes the given slice and doesn't
	  return a new one.
	*/

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings", strs)

	// NOTE: An example of sorting ints
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	// NOTE: We can also sort to check is already sorted in order
	s := sort.IntsAreSorted(ints)
	println("Sorted: ", s)
}
