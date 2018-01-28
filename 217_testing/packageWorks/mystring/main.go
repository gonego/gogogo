package main

import (
	"fmt"
	"strings"

	"github.com/gonego/gogogo/217_testing/packageWorks/mystring/mystrg"
)

const s = "Living in the fast or slow lane is completely up to you. " +
	"Only understand that life has many detours and no turning back. " +
	"Reaching your destination too soon might lose you the enjoyment" +
	"of having a beautiful view of the journey."

func main() {

	x := strings.Split(s, " ")

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystrg.Cat(x))
	fmt.Printf("\n%s\n\n", mystrg.Join(x))
}
