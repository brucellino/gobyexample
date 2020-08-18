// NOTE: https://gobyexample.com/maps

package main

import "fmt"

func main() {
	// NOTE: Maps are go's built-in associative data type. Alias dict/hash

	// NOTE: To create a map, use the builtin make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	// NOTE: The builtin len returns the number of k/v pairs
	fmt.Println("len: ", len(m))
	// NOTE: the builtin delete removes k/v pairs
	delete(m, "k2")
	fmt.Println("map: ", m)

	// NOTE: The optional second return value when getting a value from a map indicates
	// NOTE: if the key was present in the map.
	// NOTE: This can be used to disambiguate between missing keys and keys with zero values
	// NOTE: such as 0 and ""
	// NOTE: Here, we don't need the value itself, so we ignored it with the blank identifier _

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	// NOTE: Maps can also be declared and set on the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}
