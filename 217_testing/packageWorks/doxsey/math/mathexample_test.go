package math

import "fmt"

func ExampleAverage() {
	x := []float64{1, 2, 3, 4, 12, -1, -2, -3}
	fmt.Println(Average(x))
	//Output:
	//2
}

func ExampleAverage_second() {
	x := []float64{}
	fmt.Println(Average(x))
	//Output:
	//0
}

func ExampleAverage_third() {
	x := []float64{0.5, 0.6, 0.7}
	fmt.Println(Average(x))
	//Output:
	//0.6
}
