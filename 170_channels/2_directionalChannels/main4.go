
package main

import (
	"fmt"
	//"runtime"
)

func main() {
	//runtime.GOMAXPROCS(1)
	c := make(chan int)

	go send(c)

	receive(c)

	fmt.Println("about to exit")


}

// send channel
func send(c chan<- int) {
	c <- 42

	//this might not print on a multicore CPU
	fmt.Println("sent to channel")
}

// receive channel
func receive(c <-chan int) {
	fmt.Println("yet to receive anything")
	fmt.Println("the value received from the channel:", <-c)
}
