package main

import (
	"fmt"
	//"runtime"
)

func main() {
	//bidirectional channel
	c := make(chan int)

	//cs is made a channel used for sending
	cs := make(chan<- int)

	fmt.Printf("c type:\t%T\n", c)
	fmt.Printf("cs type:\t%T\n", cs)

	//a bidirectional channel can be assigned
	//to unidirectional channel without affecting their
	// original types
	cs = c
	fmt.Printf("c type after assigned to cs:\t%T\n", c)
	fmt.Printf("cs type after assigned with c:\t%T\n", cs)

	//converting c channel and assigning the result to c2
	c2 := chan<- int(c)
	fmt.Printf("c type after conversion:\t%T\n", c)
	fmt.Printf("c2 initiated and assigned as:\t%T\n", c2)
}
