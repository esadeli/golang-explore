package main

import (
	"fmt"
)

func main() {
	var x [6]int
	y := 9
	z := []int{0, 0, 0, 0} // should define the content of the array
	fmt.Println("Print array ---", x)
	fmt.Println("Print array composite literal---", z)
	fmt.Println("Print test ---", y)
	x[4] = 20
	z[2] = 7
	fmt.Println("Print array ---", x)
	fmt.Println("Print array composite literal---", z)
	fmt.Println(len(x))
	for i := 0; i < len(x); i++ {
		fmt.Println("Element of array ", x[i])
	}
}
