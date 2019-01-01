package main

import (
	"fmt"
)

const (
	a int = 20
	b     // b value will follow a
)

func main3() {
	fmt.Println(b)
	fmt.Println(b)
}
