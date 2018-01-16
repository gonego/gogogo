//func Intn(n int) int
//Intn returns, as an int, a non-negative pseudo-random
//number in [0,n) from the default Source. It panics if n <= 0.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan int)
	d := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			x := rand.Intn(12)
			c <- x
		}
		close(c)

	}()

	go func() {

		//use forever loop
		//as an alternative
		for i := 5; ; i++ {
			x := rand.Intn(12)
			d <- x
		}
	}()

	for v := range c {
		fmt.Println(v)
	}

	//control forever loop
	for i := 0; i < 5; i++ {
		fmt.Println(<-d)
	}

}
