package main

import "fmt"

// only accepts a sender channel plus message
func send(cs chan<- string, msg string) {
	cs <- msg
}

func main() {
	//this is a full channel
	//note the buffer available
	c := make(chan string, 1)

	//send full channel and message
	send(c, "gonegolang")

	//this remains a full channel
	//thus, can receive msg intact
	fmt.Println(<-c)
}
