package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//fanIn() actually needs two arguments of 
	//type chan of string receiver
	c := fanIn(boring("Joe"), boring("Ann"))
	
	//this controls the forever loop
	//in boring()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	
	//this func use a forever loop to
	//load c, a chan of string with 
	//formatted msg & i, while allowing
	//some slack period for its twin 
	//anonymous goroutine to catch-up
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	
	//this returns c, as a chan of string receiver
	return c
}

// FAN IN
// fanIn() combines both chan of string receivers
// into c, a local chan of string
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	
	//this return c as a chan of string receiver
	//to main() and assigned/iniated to a local
	//chan of string receiver
	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
