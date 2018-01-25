package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	defer foo()

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer f2.Close()
}
func foo() {
	fmt.Println("log.Panicln is called, deferred functions still run")
}

/*
log.Panicln print and a call to panic() to end current goroutine.
However dferred goroutines will still be called and will end
in reverse order
*/
