package main

import (
	"fmt"
)

func main2() {
	x4 := 20
	x5 := 30
	fmt.Println(x4 == x5) // false
	fmt.Println(x4 != x5) // true
	fmt.Println(x4 <= x5) // true
	fmt.Println(x4 >= x5) // false
	fmt.Println(x4 < x5)  // true
	fmt.Println(x4 > x5)  // false
}
