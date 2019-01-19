package main

import (
	"fmt"
)

const (
	a1 = 2018 + iota
	b1 = 2018 + iota
	c1 = 2018 + iota
	d1 = 2018 + iota
)

func main6() {
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", d1)
}
