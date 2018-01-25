package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var errMath = errors.New("error :" +
	"square root of negative number")

func main() {
	fmt.Printf("%T\n", errMath)
	x, err := sqrt(-16)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(x)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errMath
	}
	return math.Sqrt(f), nil
}
