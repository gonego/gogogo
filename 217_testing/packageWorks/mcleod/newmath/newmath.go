//Package newmath contains functions for arithmetic computation
//of sum and mean which can take an unlimited no of arguments.
package newmath

//NewAverage accept unlimited number of numerical values
//and returns the float64 value of the computed average
func NewAverage(x ...float64) float64 {
	total := 0.0
	if len(x) == 0 {
		return 0
	}
	//fmt.Printf("%T\n", x) //prints float64[]
	for _, v := range x {
		total = total + v
	}
	return total / float64(len(x))
}

//NewSum accept unlimited number of numerical values
//and returns the float64 value of the total
func NewSum(x ...float64) float64 {
	total := 0.0
	//fmt.Printf("%T\n", x) //prints float64[]
	for _, v := range x {
		total = total + v
	}
	return total
}
