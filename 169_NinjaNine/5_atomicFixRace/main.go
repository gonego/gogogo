package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	noOfRoutines := 50
	wg.Add(noOfRoutines)
	var counter int64

	fmt.Println("No of goroutines:", runtime.NumGoroutine())
	fmt.Println("Initial Counter Value:",counter)
	for i := 0; i < 50; i++ {
		go func() {
			//read, incremented and write in one go
			atomic.AddInt64(&counter,1)
			//extract current value of counter
			fmt.Println("counter:",atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("No of goroutines:", runtime.NumGoroutine())
	fmt.Println("Final Counter Value:", counter)
}
