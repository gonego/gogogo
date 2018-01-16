package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Printf("type of ctx: %T\n", ctx)
	fmt.Println("------------ERROR CHECK 1-------------------")
	fmt.Println("error 1 :", ctx.Err())
	fmt.Println("no. of goroutines:", runtime.NumGoroutine())
	fmt.Println("--------------------------------------------")

	go func() {
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				counter++
				time.Sleep(time.Millisecond * 500)
				fmt.Println("working", counter)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("------------ERROR CHECK 2-------------------")
	fmt.Println("error 2 :", ctx.Err())
	fmt.Println("no. of goroutines:", runtime.NumGoroutine())
	fmt.Println("--------------------------------------------")
	fmt.Println("cancelling ctx")
	cancel()
	fmt.Println("cancelled ctx")

	time.Sleep(time.Second)
	fmt.Println("------------ERROR CHECK 3-------------------")
	fmt.Println("error 3 :", ctx.Err())
	fmt.Println("no. of goroutines:", runtime.NumGoroutine())
}
