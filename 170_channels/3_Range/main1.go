package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go send(c)
	receive(c)

}
//send channel specific for sending
func send(c chan<- int) {
	for i := 0; i < 12; i++ {
		c <- i
	}
	close(c)
}

//send channel specific for receiving
func receive(c <-chan int) {

	// this special "range" receives
	// a sequence of v<-c
	for v := range c {
		fmt.Println(v)
	}
}
