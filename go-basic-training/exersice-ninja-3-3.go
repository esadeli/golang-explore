// Hands-on exercise #3
// Create a for loop using this syntax
// for condition { }
// Have it print out the years you have been alive.

package main

import (
	"fmt"
)

func main5() {
	x := 0
	for x <= 2000 {
		if x == 1544 {
			fmt.Println("Birth year ", x)
			break
		}
		x++
	}
}
