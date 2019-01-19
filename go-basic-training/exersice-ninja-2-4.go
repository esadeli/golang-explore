package main

import (
	"fmt"
)

func main4() {
	x := 30
	fmt.Printf("%d\t%x\t%b\n", x, x, x)

	y := x << 2
	fmt.Printf("%d\t%x\t%b\n", y, y, y)
}
