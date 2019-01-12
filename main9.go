package main

import (
	"fmt"
)

// Pointer, physical address and dereferencing
func main() {
	// pointer
	a := 21
	fmt.Println(a)
	fmt.Println(&a) // the physical address of variable a
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a               // refer b to address a
	fmt.Println(b)        // variable b store the address of a
	fmt.Println(&b)       // the address of b
	fmt.Printf("%T\n", b) // b is the pointer to variable a
	*b = 43
	fmt.Println(a)
	fmt.Println(b)          // the address of b
	fmt.Println(*b)         // dereferencing get the value stored in address of variable a
	fmt.Printf("%T\n", *&a) // variable a is pointing to its physical address
	fmt.Println(*&a)
}
