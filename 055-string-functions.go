// NOTE: https://gobyexample.com/string-functions

/*
  The standard library's strings package provides many useful string related functions.
  Here are some examples to give you a sense of the package
*/

package main

import (
	"fmt"
	s "strings"
)

// NOTE: We alias fmt.Println to be a shorter name as we'll use it a lot below
var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

	p()
	p("Len: ", len("hello"))
	p("Char: ", "hello"[1])
}
