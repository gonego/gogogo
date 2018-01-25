package main

import (
	"fmt"

	"github.com/gonego/gogogo/217_testing/packageWorks/doxsey/math"
)

func main() {
	xs := []float64{1, 2, 3}
	avg := math.Average(xs)
	fmt.Println(avg)
}
