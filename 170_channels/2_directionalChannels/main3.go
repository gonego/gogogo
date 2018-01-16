// When using channels as function parameters, you can
// specify if a channel is meant to only send or receive
// values. This specificity increases the type-safety of
// the program.
// Code adapted from https://gobyexample.com/channel-directions

package main

import "fmt"

// This `send` function only accepts a channel for sending
// values. It would be a compile-time error to try to
// receive on this channel.
func send(c1 chan<- string, msg string) {
	c1 <- msg
}

// The ` receiveNsend` function accepts one channel for receives
// (c1) and a second for sends (c2).
func receiveNsend(c1 <-chan string, c2 chan<- string) {
	msg := <-c1
	c2 <- msg
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	send(c1, "dude! this is the message")
	receiveNsend(c1, c2)
	fmt.Println(<-c2)
}
