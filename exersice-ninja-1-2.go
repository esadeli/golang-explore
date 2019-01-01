// Hands-on exercise #2
// Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
// identifier “x” type int
// identifier “y” type string
// identifier “z” type bool
// in func main
// print out the values for each identifier
// The compiler assigned values to the variables. What are these values called?

package main

import (
	"fmt"
)

var x1 int
var y1 string
var z1 bool

func main2() { // change to func main() to run
	x1 = 42
	y1 = "James Bond"
	z1 = true
	fmt.Println(x1)
	fmt.Printf("%T \n", x1)
	fmt.Println(y1)
	fmt.Printf("%T \n", y1)
	fmt.Println(z1)
	fmt.Printf("%T \n", z1)

	for i := 0; i < 3; i++ {
		fmt.Println(x1)
		fmt.Println(y1)
		fmt.Println(z1)
	}
}
