package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	noOfRoutines := 50
	counter := 0
	wg.Add(noOfRoutines)
	//runtime.GOMAXPROCS(1)
	fmt.Println("No of goroutines:", runtime.NumGoroutine())
	fmt.Println("Initial Counter Value:", counter)
	for i := 0; i < 50; i++ {
		go func() {
			x := counter
			runtime.Gosched()
			fmt.Println("routines:", runtime.NumGoroutine(),"counter:",counter)
			x++
			counter = x
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("No of goroutines:", runtime.NumGoroutine())
	fmt.Println("Final Counter Value:", counter)
}
