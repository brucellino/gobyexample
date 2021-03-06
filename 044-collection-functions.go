// NOTE: https://gobyexample.com/collection-functions

/*
We often need our programs to perform on colletions of data, like selecting all items
that satisfy a given predicate or mapping all items to a new colleciton with a
custom function.

In some languages, it's idiomatic to use generic data structures and algorithms.

Go does not support generics; in Go it's common to provide collection functions if and when they are specifically needed for your program  and data types.

Below are some collection functions for slices of strings.
Note that in some cases, it may be clearest to just inline the collection-manipulating
code directly, instead of creating and calling a helper function.
*/

package main

import (
	"fmt"
	"strings"
)

// Index returns the first index of the target string t, or -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns true of the target string t is in the slice
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns true if the target string t is in the slice.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all of the strings in the slice satisfy the predicate f
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all the strings in the slice
// that satisfy the predicate f
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying the function f
// to each of string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println("Index of pear is: ", Index(strs, "pear"))
	fmt.Println("Grape is in collection: ", Include(strs, "grape"))
	fmt.Println("Any entries with p: ", Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println("Entries with e: ", Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println("Uppercasing entries: ", Map(strs, strings.ToUpper))
}
