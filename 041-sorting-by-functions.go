// NOTE: https://gobyexample.com/sorting-by-functions
/*
Sometimes we'll want to sort a collection by something other than its natural order.
For example, suppose we wanted to sort strings by their length instead of alphabetically.
*/

package main

import (
	"fmt"
	"sort"
)

/*
In order to sort by a custom function in Go, we need a corresponding type.
We create a byLength type that is just an alias for the builtin []string type
*/
type byLength []string

/*
  We implement sort.Interface - len, less and swap - on our type, so we cna use the
  sort package's generic sort function.
  len and swap will usually be similar across types and less will hold the actual
  custom sorting logic.
  In our case we want to sort in order of increasing string length, so we use
  len(s[i]) and len(s[j]) here.
*/

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

/*
  With all of this in  place, we can now implement our custom sort by
  converting the original fruits slice to byLength, and then use sort.Sort on that
  typed slice
*/

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

/*
  By following this same pattern of creating a custom type, implemeneting the three interface
  methods on that type , and then calling sort.Sort on  collection of that custom type,
  we can sort Go slices by arbitrary functions.
*/
