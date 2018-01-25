package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		counter := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err()) //wont get printed!
				return                 //to ensure no leak
			default:
				fmt.Println(counter)
				time.Sleep(time.Millisecond * 2)
				if counter > 20 {
					return
				}
				counter++

			}
		}
	}()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Done")

}
