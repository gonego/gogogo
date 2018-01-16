//this code is copied to 170_channels/4_selectChannel/main2.go
//but with the difference in line 26
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
	done<-true //close(done)
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