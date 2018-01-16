//this is the concept of passing
//value from one channel to another
//to help understand "fan-in"

package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	d := make(chan int)

	go send(c)

	go receive(c, d)

	fmt.Println("about to exit:", <-d)

}

// send channel
func send(c chan<- int) {
	c <- 42

	//this won't print on a multicore CPU
	fmt.Println("sent to channel")
}

// receive channel
func receive(c <-chan int, d chan int) {

	x := <-c
	d <- x
}
