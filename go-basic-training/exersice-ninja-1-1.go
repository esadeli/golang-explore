// Hands-on exercise #1
// Using the short declaration operator,
// ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
// 42
// “James Bond”
// true
// Now print the values stored in those variables using
// a single print statement
// multiple print statements

package main

import (
	"fmt"
)

func main1() { // change to func main() to run
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	fmt.Println(y)
	fmt.Printf("%T \n", y)
	fmt.Println(z)
	fmt.Printf("%T \n", z)

	for i := 0; i < 3; i++ {
		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(z)
	}
}
