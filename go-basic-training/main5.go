package main

import (
	"fmt"
)

var num = 3

func main5() {
	var x [6]int // array
	y := 9
	z := []int{60, 30, 11, 9} // slices more like a dynamic array
	z2 := make([]int, 1, 2)   // create a slice
	z3 := [][]int{z, z2}      // multidiemensional slice

	// create multidimensional slices using make()
	z4 := make([][]int, 2) // multidiemensional slice
	for j := 0; j < 2; j++ {
		z4[j] = make([]int, 1)
	}
	z4[0] = z2
	z4[1] = z
	z2[0] = 1
	w := []int{2020, 2023}
	fmt.Println("Print array ---", x)
	fmt.Println("Print array ---", x[1:])
	fmt.Println("Print array ---", x[1:3])
	fmt.Println("Print slice composite literal---", z)
	fmt.Println("Print slice ---", z[1:])
	fmt.Println("Print slice ---", z[1:num])
	fmt.Println("Print test ---", y)
	x[4] = 20
	z[2] = 7
	fmt.Println("Print array ---", x)
	fmt.Println("Print slice composite literal---", z)
	fmt.Println(len(x))
	for i := 0; i < len(x); i++ {
		fmt.Println("Element of array ", x[i])
	}
	for i, v := range z {
		fmt.Println("Element of array ", i, v)
	}

	// append slice
	// can only append slice with typical content or slices with typical content
	// can't append it with normal array
	z = append(z, 88, 40)
	fmt.Println("Print slice composite literal 2---", z)
	z = append(z, w...)
	fmt.Println("Print slice composite literal 3---", z)

	// delete a component/element of a slice
	z = append(z[:1], z[2:]...)
	fmt.Println("Print slice composite literal 4---", z)

	fmt.Println("Print these ----", z2)
	z2 = append(z2, 33, 33, 55, 66, 20, 30)
	fmt.Println("Print these ----", z2)
	fmt.Println("Print these Slice.length ----", len(z2))
	fmt.Println("Print these Slice Capacity ----", cap(z2))

	// multidimensional array
	fmt.Println("multi dimensional slice ---", z3)
	fmt.Println("multi dimensional slice 2 ---", z4)

	// Map in GoLang
	num1, num2 := 10, 11 // will cause error
	fmt.Println("num1 --- ", num1)
	fmt.Println("num2 --- ", num2)

	smap1 := map[string]int{
		"Pirlo": 21,
		"Totti": 10,
	}
	fmt.Println("smap1 --- ", smap1)
	fmt.Println("Pirlo - smap 1 ", smap1["Pirlo"])
	fmt.Println("Pirlo - smap 1 ", smap1["Nesta"]) // default value of non existence parameter = 0 or empty string

	v, ok2 := smap1["Nesta"]
	fmt.Println("v --- ", v)
	fmt.Println("ok --- ", ok2)

	smap2 := map[string]string{
		"Pirlo": "Midfielder",
		"Totti": "Second Striker",
	}
	fmt.Println("smap2 --- ", smap2)
	fmt.Println("Pirlo - smap 2 ", smap2["Pirlo"])
	fmt.Println("Pirlo - smap 2 ", smap2["Nesta"])

	v2, ok := smap2["Nesta"]
	fmt.Println("v --- ", v2)
	fmt.Println("ok --- ", ok)

	// comma - ok idiom
	v3, ok := smap2["Totti"]
	fmt.Println("v --- ", v3)
	fmt.Println("ok --- ", ok)

	if v3, ok := smap2["Totti"]; ok {
		fmt.Println("This is true!", v3)
	}

	// add element
	smap2["Maldini"] = "Defender"
	fmt.Println(" smap2 updated --- ", smap2)

	// using range for MAP
	for i, k := range smap2 {
		fmt.Println("The element parameter", i)
		fmt.Println("The element content", k)
	}

	// using range for slice
	for i, k := range z4 {
		fmt.Println("The element", i)
		fmt.Println("the content", k)
	}

	// delete a key in a map
	delete(smap1, "Totti")
	fmt.Println("Updated smap1 --- ", smap1)
}
