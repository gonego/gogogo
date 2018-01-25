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

	f2, err := os.Open("secret.txt")
	if err != nil {
		log.Println("err happened...\n", err)
		fmt.Println("check the log.txt file in the directory")
	}
	defer f2.Close()

}

// Println calls Output to print to the standard logger
// which is assigned to a file instead of the stdout
