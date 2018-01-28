package main

import (
	"fmt"

	"github.com/gonego/gogogo/217_testing/packageWorks/mcleod/newmath"
)

func main() {
	x := newmath.NewAverage(1, 2, 3, 4, 5, 6, 7, 8, 9)
	y := newmath.NewSum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Average:", x)
	fmt.Println("Total:", y)

}
