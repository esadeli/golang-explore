package main

import "fmt"

func main() {
	ch := make(chan int)
	go send(ch)
	receive(ch)
	// close(ch) // will close the channel without executing the next line of code
	fmt.Println("About to exit")
}

// send
func send(c chan<- int) {
	// c <- 50
	// if using range / loop
	for i := 0; i < 25; i++ {
		c <- i
	}
	close(c) // range will keep looping until the channel close
}

// receive
func receive(c <-chan int) {
	for v := range c {
		fmt.Println("The value received from channel is ", v)
	}
	// fmt.Println("The value received from channel is ", <-c) // c will only give physical address
}
