package main

import (
	"fmt"
)

// var x int

func main() {
	a := add()
	b := add()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	// factorial
	sumfac := factorial(7)
	fmt.Println(sumfac)
}

// closure for return function
func add() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// recursion
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num + factorial(num-1)
}
