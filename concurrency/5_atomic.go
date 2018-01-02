package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64 = 0

	const gs = 60
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter value", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	atomic.LoadInt64(&counter)
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
