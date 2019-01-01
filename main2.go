package main

import (
	"fmt"
	"runtime"
)

// const
const (
	a int    = 50
	b string = "Constant string"
	c bool   = true
)

// const - iota
const (
	d = iota
	e
	f
)

// boolean
var a1 bool
var num1 int
var num2 float64

func main7() {
	fmt.Println("hello")
	fmt.Println(a1)
	a1 = true
	fmt.Println(a1)

	// operator comparison
	a2 := 40
	a3 := 2
	fmt.Println(a2 == a3) // false
	fmt.Println(a2 != a3) // true
	fmt.Println(a2 >= a3) // true

	// numeric types
	num1 = 40
	num2 = 50.0231231
	fmt.Println(num1)
	fmt.Printf("%T\n", num1)
	fmt.Println(num2)
	fmt.Printf("%T\n", num2)

	// runtime package
	fmt.Println(runtime.GOOS)   // get the operating system
	fmt.Println(runtime.GOARCH) // get the program's architecure target

	// string variable
	str := "Golang lecture, 世界"
	fmt.Println(str)
	fmt.Printf("%s\n", str)
	fmt.Printf("%q\n", str)      // "Golang lecture" // quoted verb
	fmt.Printf("%x\n", str)      // 476f6c616e67206c656374757265 // hexadecimal
	fmt.Printf("---- %x\n", "世") // a symbol of e4b896

	for k := 0; k < len(str); k++ {
		fmt.Printf("---%x\t", str[k])
	}

	// string to byte conversion
	bs := []byte(str)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// utf-8
	for l := 0; l < len(str); l++ {
		fmt.Printf("--- %U", str[l])
	}

	for l, v := range str {
		// fmt.Println(l, v)
		fmt.Printf("at index position %d we have hex %#x\n", l, v)
	}

	// const in GoLang
	fmt.Printf("The variable value is %d\n", a)
	fmt.Printf("The variable value is %s\n", b)
	fmt.Printf("The variable value is %t\n", c)

	// const iota in GoLang
	fmt.Println("IOTA", d)
	fmt.Println("IOTA", e)
	fmt.Println("IOTA", f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)

	// bit shifting
	x := 4
	fmt.Printf("%d\t\t\t%b\n", x, x)
	y := x << 2 // multiply 4 by 2 and another 2
	fmt.Printf("%d\t\t%b\n", y, y)

	// bit shifting 2
	kb := 1024
	mb := kb * 1024
	gb := mb * 1024
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
}
