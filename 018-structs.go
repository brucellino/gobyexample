package main

import "fmt"

// NOTE: Go's structs are typed collections of of fields.
// NOTE: They are useful for grouping data together to form records

type person struct {
	// NOTE: This person struct type has name and age fields
	name string
	age  int
}

func newPerson(name string) *person {
	// NOTE: newPeron constructs a new peron struct with the given name
	p := person{name: name}
	p.age = 42
	// NOTE: You can return a pointer to a local variable, since a local variable will survive
	// NOTE: the scope of the function
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})              // NOTE: Create a new person struct
	fmt.Println(person{name: "Alice", age: 30}) // NOTE: can name fields when initialising
	fmt.Println(person{name: "Fred"})           // NOTE: omitted fields will be zero-valued
	fmt.Println(&person{name: "Ann", age: 40})  // NOTE: an & prefix yeilds a pointer to the struct
	fmt.Println(newPerson("Jon"))               // NOTE: it's idiomatic to encapsulate new struct creation in constructor functions

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // NOTE: you can access fields with a dot

	sp := &s
	fmt.Println(sp.age) // NOTE: you can also use dots with struct pointers
	// NOTE: the pointers are automatically dereferenced

	sp.age = 51 // NOTE: structs are mutable
	fmt.Println(sp.age)
}
