package main

import (
	"fmt"
	"log"
	"math"
)

type mathError struct {
	part   string
	number int
	err    error
}

func (n mathError) Error() string {
	return fmt.Sprintf("\nPart:%v  No:%v Error:%v", n.part, n.number, n.err)
}

func main() {
	x, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(x)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		errorMsg := fmt.Errorf("square root of negative number: %v", f)
		return 0, mathError{"A", 25, errorMsg}
	}
	return math.Sqrt(f), nil
}
