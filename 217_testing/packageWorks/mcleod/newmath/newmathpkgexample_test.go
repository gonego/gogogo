package newmath_test

import (
	"fmt"

	"github.com/gonego/gogogo/217_testing/packageWorks/mcleod/newmath"
)

type test struct {
	value []float64
}

func Example() {

	a := []test{
		test{[]float64{1, 2, 3}},
		test{[]float64{-1, 1}},
		test{[]float64{10, 20, 30}},
		test{[]float64{}},
	}
	for _, v := range a {
		fmt.Println(newmath.NewAverage(v.value...))
	}
	//Output:
	//2
	//0
	//20
	//0
}
