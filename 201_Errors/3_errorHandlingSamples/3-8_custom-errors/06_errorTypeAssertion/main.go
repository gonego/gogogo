package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("we have problem: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more lemonade",
	}

	bar(c1)
}

func bar(e error) {
	fmt.Println("msg from bar -", e)
	fmt.Println(e.(customErr).info, "!!!")
}
