package main

import (
	"fmt"
)

func main() {

	//cs is made a channel used for sending
	cs := make(chan<- int)

	go func() {
		//42 is sent thru channel cs
		//waiting to be received
		cs <- 42

	}()

	fmt.Printf("cs type:%T\n", cs) // cs type: chan<-int

	//cs by design is not a channel for receiving
	//error message if we try to run this line:
	//fmt.Println(<-cs)

}
