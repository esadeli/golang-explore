package main

import (
	"fmt"
)

type car struct {
	brand  string
	origin string
}

type person struct {
	name    string
	country string
}

type thing interface {
	shoutout()
}

func main7() {
	defer deferExample()
	sayHello("Andara")
	s1 := sayMorning("Gunawan")
	fmt.Println(s1)

	// get sum
	arrNum := []int{2, 3, 55, 4, 10, 8}
	result := calculateSum("Andi", 2, 3, 55, 4, 10, 8)
	// result := calculateSum(arrNum) // if stated as an array of integer
	fmt.Println(result)
	result2 := calculateSum("Budi", arrNum...)
	fmt.Println(result2)

	result3 := calculateSum("") // you still have to assign the string value
	fmt.Println(result3)

	car1 := car{
		brand:  "BMW",
		origin: "Germany",
	}
	car1.shoutout()

	person1 := person{
		name:    "Sarah",
		country: "United States",
	}
	person1.shoutout()

	// polymorphism using interface
	shoutGeneral(car1)
	shoutGeneral(person1)

	// anonymous function
	func(x int) {
		fmt.Println("Return ", x+2)
	}(5)

	// func expression
	expF := func(s string) {
		fmt.Println("I print whatever is given ", s)
	}
	expF("Merdeka!!!")

	// return function
	retFunc1 := retFunc(30)
	fmt.Println("return function", retFunc1)     // 0x487d60
	fmt.Println("return function 2", retFunc1()) // 0
	fmt.Printf("%T\n", retFunc1)
	fmt.Printf("%T\n", retFunc1()) // int

	retFunc2 := retFunc1()
	fmt.Println("return function 2", retFunc2) // 30

	fmt.Println(retFunc(40)())

	// callback in go
	testCbsum := evenSum(sumCb, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	fmt.Println("Callback test --- ", testCbsum)
}

func sayHello(name string) {
	fmt.Println("Hello ", name)
}

func sayMorning(name string) string {
	return fmt.Sprint("Morning ", name)
}

func calculateSum(s string, num ...int) int { //variadic parameter should be put as the last parameter
	// func calculateSum(num []int) int { // if stated as an array of integer
	fmt.Println("Hi %s", s)
	fmt.Printf("%T\n", num)
	sum := 0
	for i, v := range num {
		fmt.Println("Element ", i, " with sum ", sum)
		sum += v
	}
	return sum
}

func deferExample() {
	fmt.Println("Put first and show last")
}

// method --> function that is attached to a struct
func (data car) shoutout() {
	fmt.Println("Hey this car is a ", data.brand, " and it is from ", data.origin)
}

// polymorphism of shoutout function
func (data person) shoutout() {
	fmt.Println("Hey my name is ", data.name, " and I am from ", data.country)
}

// polymorphism using interface
func shoutGeneral(data thing) {
	switch data.(type) {
	case car:
		fmt.Println("Mobil ini bermerek ", data.(car).brand, " dan berasal dari ", data.(car).origin)
	case person:
		fmt.Println("Hai nama saya ", data.(person).name, " saya berasal dari ", data.(person).country)
	}
}

// function returning a function
func retFunc(angka int) func() int {
	return func() int {
		return angka
	}
}

// callbacks
func sumCb(angka ...int) int {
	sum := 0
	for _, v := range angka {
		sum += v
	}
	return sum
}

func evenSum(f func(angka ...int) int, numbers ...int) int {
	numArr := []int{}

	for _, k := range numbers {
		if k%2 != 0 {
			numArr = append(numArr, k)
		}
	}
	fmt.Println(numArr)
	Total := f(numArr...)
	return Total
}
