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

	f2, err := os.Open("no-file.txt")

	//this wont be called at all
	//thus may be omitted
	defer f2.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

/*
log.Fatalln functions call os.Exit(1) after writing the log message
*/
