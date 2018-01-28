package math_test

import (
	"fmt"

	"github.com/gonego/gogogo/217_testing/packageWorks/doxsey/math"
)

var xs []float64

func Example() {
	xs := []float64{1, 1, 1}
	avg := math.Average(xs)
	fmt.Println(avg)
	//Output:
	//1
}
