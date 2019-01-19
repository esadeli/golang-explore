package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mu sync.Mutex

func main11c() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	wg.Add(2)
	go incrementor("Foo")
	go incrementor("bar")
	wg.Wait()
	fmt.Println("Final counter: ", counter)
}

func incrementor(st string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		// fmt.Println("index ", i)
		mu.Lock()
		counter++
		fmt.Println(st, i, "Counter: ", counter)
		mu.Unlock()
	}
	wg.Done()
}
