// Race condition

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var mu sync.Mutex

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines 1: ", runtime.NumGoroutine())

	// counter := 0
	var counter int64

	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// mu.Lock()
			// v := counter
			// runtime.Gosched()
			// v++
			// counter = v
			// fmt.Println("Counter 1: ", counter)
			// mu.Unlock()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter 1 ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("GoRoutines 2: ", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("GoRoutines 3: ", runtime.NumGoroutine())
	fmt.Println("Counter 2: ", counter)
}
