
package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 60
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			//read the value
			v := counter

			//one alternative:
			//func Sleep(d Duration)
			//Sleep pauses the current goroutine
			//for at least the duration d
			//time.Sleep(time.Second)

			//func Gosched()
			//Gosched yields the processor, allowing
			//other goroutines to run. It does not
			//suspend the current goroutine,
			//so execution resumes automatically
			//runtime.Gosched()

			//increment the value
			v++

			//write back the value
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
