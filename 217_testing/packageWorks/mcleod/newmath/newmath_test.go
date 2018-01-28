package newmath_test

import (
	"testing"

	"github.com/gonego/gogogo/217_testing/packageWorks/mcleod/newmath"
)

type testSet struct {
	value  []float64
	result float64
}

func TestNewAverage(t *testing.T) {

	a := []testSet{
		testSet{[]float64{1, 2, 3}, 2},
		testSet{[]float64{-1, 1}, 0},
		testSet{[]float64{10, 20, 30}, 20},
		testSet{[]float64{}, 0},
	}
	for _, v := range a {
		x := newmath.NewAverage(v.value...)
		if x != v.result {
			t.Error("For:", v.value, "Expected:", v.result, "Got:", x)
		}
	}
}

func TestNewSum(t *testing.T) {
	a := []testSet{
		testSet{[]float64{1, 2, 3}, 6},
		testSet{[]float64{-1, 1}, 0},
		testSet{[]float64{10, 20, 30}, 60},
		testSet{[]float64{}, 0},
	}
	for _, v := range a {
		x := newmath.NewSum(v.value...)
		if x != v.result {
			t.Error("For:", v.value, "Expected:", v.result, "Got:", x)
		}
	}
}
