// NOTE: https://gobyexample.com/string-formatting
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	/*
	  Go offers several printing "verbs" designed to format general Go values.
	  For example, this prints an instance of our point struct.
	*/
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	// NOTE: If the value is a struct, the %+v variant will include the struct's field names
	fmt.Printf("%+v\n", p)
	// NOTE: The %#v variant prints a Go syntax representation of the value, ie the source
	// NOTE: code snippet that would produce that value
	fmt.Printf("%#v\n", p)
	// NOTE: To print the type of the value, use %T
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)                         // printing booleans
	fmt.Printf("%d\n", 123)                          // integers
	fmt.Printf("%b\n", 14)                           // binary
	fmt.Printf("%c\n", 33)                           // character
	fmt.Printf("%x\n", 456)                          // hex
	fmt.Printf("%f\n", 78.9)                         // floats
	fmt.Printf("%e\n%E\n", 123400000.0, 123400000.0) // scientific
	fmt.Printf("%s\n", "\"string\"")                 // string
	fmt.Printf("%p\n", &p)                           // pointer
	fmt.Printf("%q\n", "\"string\"")                 // double-quoted strings
	fmt.Printf("|%6.2f|%6f|\n", 1.2, 3.45)           // Control width and precision of numbers\
	fmt.Printf("|%-6.2s|%6.2s|\n", "foo", "b")       // left-justify with - flag

	s := fmt.Sprintf("a %s", "string") // formatted print to a string
	fmt.Println(s)

	fmt.Fprint(os.Stderr, "an %s\n", "error")
}
