//Package math contains simple math function
//used in training of Go package manipulation
package math

//Average find the mean of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return 0
	}
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))

}

/*
...must run "go install" within this
...folder dir before running other file
...that uses this package
*/
