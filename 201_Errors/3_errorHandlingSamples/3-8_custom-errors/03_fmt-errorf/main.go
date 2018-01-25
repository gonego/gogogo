package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	x, err := sqrt(-25)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(x)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("error again:"+
			"square root of negative number: %v", f)
	}
	return math.Sqrt(f), nil
}
