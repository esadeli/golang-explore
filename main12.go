package main

import "fmt"

func main12() {
	// c := make(chan int) // without buffer channel
	c := make(chan int, 2)
	// ch := make(chan<- int, 1) // a channel that only receive string variable
	// you can only run and use channel using go concurrency

	// without buffer channel
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		c <- 12
	// 	}
	// }()

	// // fmt.Println(<-c)
	// go func() {
	// 	for {
	// 		fmt.Println(<-c)
	// 	}
	// }()
	// time.Sleep(time.Second)

	// with buffer channel
	// a channel should have its own in and out,
	c <- 33
	c <- 34 // will create error if we don't let the value out by creating buffer channel
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
}
