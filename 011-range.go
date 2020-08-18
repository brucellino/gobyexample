// NOTE: https://gobyexample.com/range

// NOTE: Range iterates over elements in a variety of data structures.

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// NOTE: Range on arrays and slices provide both the index and value for each entry
	// NOTE: We don't need the index, so we ignore it with _
	// NOTE: Sometimes we do want the identifier though
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// NOTE: range can also iterate over just the keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// NOTE: Range on strings iterates over unicode code points
	// NOTE: The first value is the starting byte inde, of the rune
	// NOTE: and the second is the rune itself
	for i, c := range "🦟" {
		fmt.Println(i, c)
	}
}
