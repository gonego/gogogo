package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Printf("Type of ctx: %T\n", ctx)
	fmt.Println("---But this type is unexported---\n")

	//not until function main completes
	defer cancel()

	for v := range getNumbers(ctx) {
		fmt.Println(v)
		if v == 10 {
			break
		}
	}
}

func getNumbers(ctx context.Context) <-chan int {
	numbers := make(chan int)
	counter := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return //to ensure no leak
			case numbers <- counter:
				counter++
			}
		}
	}()
	return numbers
}
