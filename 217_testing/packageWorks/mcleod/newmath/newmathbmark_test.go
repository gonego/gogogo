package newmath_test

import (
	"testing"

	"github.com/gonego/gogogo/217_testing/packageWorks/mcleod/newmath"
)

func BenchmarkNewAverage(b *testing.B) {

	for i := 0; i < b.N; i++ {
		newmath.NewAverage(1, 2, 3, 4, 5, 6, 7, 8, 9)

	}
}

func BenchmarkNewSum(b *testing.B) {

	for i := 0; i < b.N; i++ {
		newmath.NewSum(1, 2, 3, 4, 5, 6, 7, 8, 9)

	}
}
