//this code is similar to 170_channels/4_selectChannel/main.go
//except for line 26 where close(done) replaces done<-true
package main

import "fmt"

func main(){
	even:=make(chan int)
	odd:=make(chan int)
	done:=make(chan bool)

	go send(even,odd,done)

	receiver(even,odd,done)
	fmt.Println("Done...")
}

func send(even chan<- int,odd chan<- int, done chan<- bool){
	for i:=1;i<100;i++{
		if i%2==0{
			even <-i
		}else{
			odd<-i
		}
	}
	close(done)
}

func receiver(even <-chan int,odd <-chan int, done <-chan bool){
	for{
		select {
		case v := <-even:
			fmt.Println("Even:", v)
		case v := <-odd:
			fmt.Println("Odd:", v)
		case <-done:
			return
		}
	}
}