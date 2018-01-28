//Package mystrg contains 2 functions used for a similar purpose
//of concatenating the elements of a slice of strings to create
//a single string. A separator string will be placed between
//elements in the resulting string.
package mystrg

import "strings"

//Cat takes a slice of strings and join all the element into
//a single string by means of customized algorithm.
//Space will be placed between the substrings
func Cat(x []string) string {
	s := x[0]
	for _, v := range x[1:] {
		s += " "
		s += v
	}
	return s
}

//Join takes a slice of strings and join all the element into
//a single string by means of function strings.Join.
//Space will be placed between the substrings
func Join(x []string) string {
	return strings.Join(x, " ")
}
