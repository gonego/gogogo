package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	fmt.Printf("%p\n",c)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	//this "range" version will receive
	// as if v<-c
	for v:= range c {
		fmt.Println(v)
	}

	//assign a totally new channel
	c=make(chan int)

	//will print different address
	fmt.Printf("%p\n",c)

	go func(){
		c<-100
	}()

	//,ok idiom can be used for testing
	// whether channel is open or closed
	v,ok:=<-c
	if ok==false{
		fmt.Println("Channel closed")
	}else{
		fmt.Println(v)
	}



}
