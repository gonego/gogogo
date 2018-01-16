package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)
	d := make(chan int)
	e := make(chan int)
	fanin := make(chan int)

	go sender(c, d, e)
	go receiver(c, d, e, fanin)

	for v := range fanin {
		fmt.Println("Value:", v)
	}

	fmt.Println("About to exit...")

}

// send channel
func sender(c, d, e chan<- int) {

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		wg.Done()
	}()
	go func() {
		for i := 5; i < 10; i++ {
			d <- i
		}
		wg.Done()

	}()
	go func() {
		for i := 10; i < 15; i++ {
			e <- i
		}
		wg.Done()

	}()

	wg.Wait()
	close(c)
	close(d)
	close(e)
}

//receive channel
func receiver(c, d, e <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		for v := range c {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range d {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range e {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)

}
