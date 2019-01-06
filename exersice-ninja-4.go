// Hands-on exercise #1
// Using a COMPOSITE LITERAL:
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// print out the TYPE of the array

// Hands-on exercise #2
// Using a COMPOSITE LITERAL:
// create a SLICE of TYPE int
// assign 10 VALUES
// Range over the slice and print the values out.
// Using format printing
// print out the TYPE of the slice

// Hands-on exercise #3
// Using the code from the previous example,
// use SLICING to create the following new slices which are then printed:
// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]

// Hands-on exercise #4
// Follow these steps:
// start with this slice
// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// append to that slice this value
// 52
// print out the slice
// in ONE STATEMENT append to that slice these values
// 53
// 54
// 55
// print out the slice
// append to the slice this slice
// y := []int{56, 57, 58, 59, 60}
// print out the slice

// Hands-on exercise #5
// To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
// start with this slice
// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
// [42, 43, 44, 48, 49, 50, 51]

// Hands-on exercise #6
// Create a slice to store the names of all of the states in the United States of America.
// What is the length of your slice? What is the capacity? Print out all of the values,
// along with their index position in the slice, without using the range clause.
// Here is a list of the states:

// ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`,
// ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`,
// ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`,
// ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`,
// ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`,
// ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
// ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`,
// ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`,
// ` Wisconsin`, ` Wyoming`,

// Hands-on exercise #7
// Create a slice of a slice of string ([][]string).
// Store the following data in the multi-dimensional slice:

// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."

// Range over the records, then range over the data in each record.

// Hands-on exercise #8
// Create a map with a key of TYPE string which is a person’s “last_first” name,
// and a value of TYPE []string which stores their favorite things.
// Store three records in your map. Print out all of the values,
// along with their index position in the slice.

//     `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
// `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
// `no_dr`, `Being evil`, `Ice cream`, `Sunsets`

// Hands-on exercise #9
// Using the code from the previous example, add a record to your map.
// Now print the map out using the “range” loop

// Hands-on exercise #10
// Using the code from the previous example, delete a record from your map.
// Now print the map out using the “range” loop

package main

import (
	"fmt"
)

func main() {
	// exersice 1
	ex1 := [5]int{10, 2, 5, 72, 2}

	for i, k := range ex1 {
		fmt.Printf("%T\n", k)
		fmt.Println("value ex1 --- ", i, k)
	}
	fmt.Printf("ex 1 type %T\n", ex1)

	// exersice 2
	ex2 := []int{22, 45, 20, 3, 6, 7, 9, 10, 12, 23}

	for i, k := range ex2 {
		fmt.Printf("%T\n", k)
		fmt.Println("value ex2--- ", i, k)
	}
	fmt.Printf("ex2 type %T\n", ex2)

	// exersice 3
	ex3 := make([]int, 10)
	for i := 0; i < len(ex3); i++ {
		ex3[i] = 42 + i
	}

	fmt.Println(ex3)
	fmt.Println(ex3[:5])  // [42 43 44 45 46]
	fmt.Println(ex3[5:])  // [47 48 49 50 51]
	fmt.Println(ex3[2:7]) // [44 45 46 47 48]
	fmt.Println(ex3[1:6]) // [43 44 45 46 47]

	// exersice 4
	ex4 := make([]int, 10)

	for i := 0; i < len(ex4); i++ {
		ex4[i] = 42 + i
	}
	fmt.Println(ex4)

	ex4 = append(ex4, 52)
	fmt.Println(ex4)
	ex4 = append(ex4, 53, 54, 55)
	fmt.Println(ex4)

	ex4a := []int{56, 57, 58, 59, 60}
	fmt.Println(ex4a)

	ex4 = append(ex4, ex4a...)
	fmt.Println(ex4)

	// exersice 5
	ex5 := make([]int, 10)

	for i := 0; i < len(ex5); i++ {
		ex5[i] = 42 + i
	}
	fmt.Println(ex5)

	ex5a := append(ex5[:3], ex5[6:]...)
	fmt.Println(ex5a)

	// exersice 6
	ex6 := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	for i := 0; i < len(ex6); i++ {
		fmt.Println(ex6[i])
	}

	// exersice 7
	ex7a := []string{"James", "Bond", "Shaken, not stirred"}
	ex7b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	ex7 := [][]string{ex7a, ex7b}

	fmt.Println(ex7)

	// exersice 8
	ex8a := []string{`Shaken, not stirred`, `Martinis`, `Women`}
	ex8b := []string{`James Bond`, `Literature`, `Computer Science`}
	ex8c := []string{`Being evil`, `Ice cream`, `Sunsets`}
	ex8 := map[string][]string{
		"bond_james":      ex8a,
		"moneypenny_miss": ex8b,
		"no_dr":           ex8c,
	}

	for i, j := range ex8 {
		fmt.Println("This is element ", i)
		for k, l := range j {
			fmt.Println("The value of ", k, " is ", l)
		}
	}

	// exersice 9
	ex9 := []string{`Programming`, `Music`, `Book`}
	ex8["darry"] = ex9 // this object will not  be at the latest parameter if we used Capital letter

	for i, j := range ex8 {
		fmt.Println("This is element", i)
		for k, l := range j {
			fmt.Println("The value of ", k, " is ", l)
		}
	}

	// exersice 10
	delete(ex8, "bond_james")
	for i, j := range ex8 {
		fmt.Println("This is element", i)
		for k, l := range j {
			fmt.Println("The value of ", k, " is ", l)
		}
	}
}
