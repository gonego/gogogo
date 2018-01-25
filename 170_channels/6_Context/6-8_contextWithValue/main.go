package main

import (
	"context"
	"fmt"
)

type ctxKey string

func main() {
	k := ctxKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	fmt.Println(ctx)
	foo(ctx, k)

	ctx = context.WithValue(ctx, ctxKey("database"), "mysql")
	fmt.Println(ctx)
	foo(ctx, ctxKey("database"))

	foo(ctx, ctxKey("os"))
}

func foo(ctx context.Context, s ctxKey) {
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", s)
}
