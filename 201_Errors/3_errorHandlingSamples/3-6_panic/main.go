package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}

func foo() {
	fmt.Println("this is foo")
}

// http://godoc.org/builtin#panic
