package math

import "fmt"

func ExampleAverage() {
	x := []float64{1, 2, 3, 4, 12, -1, -2, -3}
	fmt.Println(Average(x))
	//Output:
	//2
}
