package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup
	noOfRoutines := 50
	counter := 0
	wg.Add(noOfRoutines)
	//runtime.GOMAXPROCS(1)
	fmt.Println("No of goroutines:", runtime.NumGoroutine())
	fmt.Println("Initial Counter Value:", counter)
	for i := 0; i < 50; i++ {
		go func() {
			//lock the area where race condition starts
			m.Lock()

			x := counter
			//runtime.Gosched()
			// does not make sense since we have Mutex lock
			fmt.Println("routines:", runtime.NumGoroutine(),"counter:",counter)
			x++
			counter = x

			//unlock once race condition prevented
			m.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("No of goroutines:", runtime.NumGoroutine())
	fmt.Println("Final Counter Value:", counter)
}
