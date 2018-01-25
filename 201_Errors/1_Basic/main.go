package main

import (
	"fmt"
	"time"
)

// MyError is an error implementation that includes
// a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Time: %v\nError: %v\n",
		e.When.Format("02-Jan-2006 03:04:05 pm Mon"), e.What)
}

func oops() error {
	return MyError{
		When: time.Now(),
		What: "Oops I did it again",
	}
}

func main() {
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().Zone())
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
