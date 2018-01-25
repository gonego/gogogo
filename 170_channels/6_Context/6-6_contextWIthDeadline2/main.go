package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	go func() {

		for {
			select {
			case <-time.After(40 * time.Millisecond):
				fmt.Println("overslept")
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			}
		}
	}()
	time.Sleep(150 * time.Millisecond)
}
