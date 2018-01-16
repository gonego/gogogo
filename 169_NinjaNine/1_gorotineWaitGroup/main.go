package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//limit to 2 cores just to reduce interleaving of print output
	runtime.GOMAXPROCS(2)
	fmt.Println("No of goroutines:", runtime.NumGoroutine())
	wg.Add(2)
	go func() {
		fmt.Println("No of goroutines:", runtime.NumGoroutine())
		for i := 0; i < 10; i++ {
			fmt.Println("Buzz:", i)
		}
		wg.Done()
	}()
	go func() {
		fmt.Println("No of goroutines:", runtime.NumGoroutine())
		for i := 0; i < 10; i++ {
			fmt.Println("Fizz:", i)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("No of goroutines:", runtime.NumGoroutine())
}
