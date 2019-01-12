// Hands-on exercise #2
// create a func with the identifier foo that
// takes in a variadic parameter of type int
// pass in a value of type []int into your func (unfurl the []int)
// returns the sum of all values of type int passed in
// create a func with the identifier bar that
// takes in a parameter of type []int
// returns the sum of all values of type int passed in

// Hands-on exercise #3
// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

// Hands-on exercise #4
// Create a user defined struct with
// the identifier “person”
// the fields:
// first
// last
// age
// attach a method to type person with
// the identifier “speak”
// the method should have the person say their name and age
// create a value of type person
// call the method from the value of type person

// Hands-on exercise #5
// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

// Hands-on exercise #6
// Build and use an anonymous func

// Hands-on exercise #7
// Assign a func to a variable, then call that func

// Hands-on exercise #8
// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

// Hands-on exercise #9
// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument

// Hands-on exercise #10
// Closure is when we have “enclosed” the scope of a variable in some code block.
// For this hands-on exercise, create a func which “encloses” the scope of a variable:

package main

import (
	"fmt"
	"math"
)

type person struct {
	first string
	last  string
	age   int
}

type square struct {
	length float64
	width  float64
	area   float64
}

type circle struct {
	radius float64
	area   float64
}

type shape interface {
	calcArea() float64
}

func main() {
	// Exersice 3
	defer exampleDefer()

	// Exersice 3
	result := foo([]int{2, 3, 4}...)
	fmt.Println("hasil --- ", result)

	result2 := foo2([]int{1, 2, 3, 4, 5})
	fmt.Println("Hasil 2 --- ", result2)

	// Exersice 4
	person1 := person{
		first: "Andrea",
		last:  "Pirlo",
		age:   40,
	}
	person1.speak()

	// Exersice 5
	circle1 := circle{
		radius: 2,
	}
	// circle1.calcArea() // --> alternative working solution
	info(circle1)

	square1 := square{
		length: 4.5,
		width:  3,
	}
	info(square1)
	// square1.calcArea() // --> alternative working solution

	// Exersice 6
	func(num1 int, num2 int) {
		fmt.Println("The result of a + b is ", num1+num2)
	}(3, 5)

	// Exersice 7
	sqrFunc := func(num int) int {
		// fmt.Println("The square of ", num, " is ", num*num)
		return num * num
	}
	// sqrFunc(3) // ---> alternative working solution
	fmt.Println("The square is ", sqrFunc(3))

	// Exersice 8
	retFunc1 := retFunc(3, 4)
	retFunc1()
	fmt.Println("Return function ----- ", retFunc1())

	// Exersice 9
	fmt.Println("Callback calculation ", cbFunc(3, getCb))
}

// Exersice 2
func foo(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func foo2(num []int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

// Exersice 3
func exampleDefer() {
	fmt.Println("Put first but print last")
}

// Exersice 4
func (data person) speak() {
	fmt.Println("This is person name is ", data.first, " ", data.last, " He/She is ", data.age, " year(s) old")
}

// Exersice 5
func (area circle) calcArea() float64 {
	return math.Pi * area.radius * area.radius
}

func (area square) calcArea() float64 {
	return area.length * area.width
}

func info(s shape) {
	fmt.Println("The area is ", s.calcArea())
}

// Exersice 8
func retFunc(num1 int, num2 int) func() int {
	return func() int {
		return num1 * num2 * 3
	}
}

// Exersice 9
func cbFunc(num1 int, f func(num2 int) int) int {
	return num1 + f(num1)
}

func getCb(num1 int) int {
	return num1 * 2
}
