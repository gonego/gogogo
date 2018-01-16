package main

import (
	"context"
	"fmt"
)

func main() {

	//func Background() Context
	//Background returns a non-nil, empty Context. It is never
	//canceled, has no values, and has no deadline. It is typically
	//used by the main function, initialization, and tests, and as
	//the top-level Context for incoming requests.
	ctx := context.Background()

	//Canceled is the error returned by Context.Err when the
	//context is canceled. DeadlineExceeded is the error returned
	//by Context.Err when the context's deadline passes.
	fmt.Println("context:", ctx)
	fmt.Println("context err:", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("----------------------------")

	//func WithCancel(parent Context) (ctx Context, cancel  CancelFunc)
	//WithCancel returns a copy of parent with a new Done
	//channel. The returned context's Done channel is closed when
	//the returned cancel function is called or when the parent
	//context's Done channel is closed, whichever happens first.
	//Canceling this context releases resources associated with it,
	//so code should call cancel as soon as the operations running
	//in this Context complete.
	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:", ctx)
	fmt.Println("context err:", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("cancel:", cancel)
	fmt.Printf("cancel type: %T\n", cancel)
	fmt.Println("----------------------------")

	cancel()

	fmt.Println("context:", ctx)
	fmt.Println("context err:", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("cancel:", cancel)
	fmt.Printf("cancel type: %T\n", cancel)
	fmt.Println("----------------------------")

}
