package main

import (
	"fmt"
)

func main() {
	//this will not work if chan resides
	//within the same goroutine, unless
	//a buffer is created:
	/*
		c:=make(chan int)
		c<-69
		fmt.Println(<-c)
	*/

	//an unsynchronized buffer of
	//2 values to send it later
	c1 := make(chan int, 2)
	c1 <- 20
	c1 <- 18
	fmt.Printf("I prevail in %v%v\n", <-c1, <-c1)

	c2 := make(chan string)
	go func() {
		//hooks c2 to func()
		c2 <- "this is unbuffered, but synchronized"
	}()
	fmt.Println(<-c2)

	go func() {
		//hooks c2 to func()
		doSomething(c2)
	}()
	<-c2 //value "done" will be discarded

}

//cx and c2 refer to the same channel, must wait
//until c2 is available from prior engagement
func doSomething(cx chan string) {
	fmt.Println("done and over")
	cx <- "done"
}
