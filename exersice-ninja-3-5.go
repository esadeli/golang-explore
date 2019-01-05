// Hands-on exercise #6
// Create a program that shows the “if statement” in action.

// Hands-on exercise #7
// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

// Hands-on exercise #8
// Create a program that uses a switch statement with no switch expression specified.

package main

import (
	"fmt"
)

func main5() {
	x := 6
	if x > 0 {
		switch {
		case (x%2 == 0 && x%3 != 0):
			fmt.Println("Even number ---", x)
		case (x%3 == 0):
			fmt.Println("Number divisable by 3 ---", x)
		default:
			fmt.Println("Default ---", x)
		}
	} else {
		fmt.Println("Negative number")
	}
}
