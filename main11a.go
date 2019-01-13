package main

import (
	"fmt"
	"runtime"
	"sync"
)

// concurrent function result won't show on the terminal
// because the system only show the result on the main thread
// therefore we have to "add information" to the system that we have another (n) thread using WaitGroup
var waitGrp sync.WaitGroup

func main11a() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\n", runtime.NumGoroutine())

	// concurrency
	waitGrp.Add(1)
	go foo() // how to create concurency in Go
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\n", runtime.NumGoroutine())
	otherFoo()
	waitGrp.Wait()

	// foo channel
	ch := make(chan int)
	go func() {
		ch <- fooChannel(4)
	}()
	fmt.Println(<-ch) // without this we can't get the return of other function
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Show Foo", i)
	}
	waitGrp.Done()
}

func otherFoo() {
	for i := 0; i < 8; i++ {
		fmt.Println("Show Other Foo", i)
	}
}

func fooChannel(num int) int {
	for i := 0; i < num; i++ {
		fmt.Println("Show Foo Channel", i)
	}
	return 0 // can't return empty return
}
