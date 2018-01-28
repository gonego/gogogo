package newmath_test

import (
	"fmt"

	"github.com/gonego/gogogo/217_testing/packageWorks/mcleod/newmath"
)

func ExampleNewAverage() {
	fmt.Println(newmath.NewAverage(1, 2, 3))
	//Output:
	//2

}

func ExampleNewSum() {
	fmt.Println(newmath.NewSum(1, 2, 3))
	//Output:
	//6
}
