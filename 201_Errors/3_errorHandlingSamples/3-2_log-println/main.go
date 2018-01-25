// log.Println outputs to the standard logger
// either to a specified file or by default to stdout
// the screen output is similar to fmt.Println but
// includes the date-time stamp

package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("secret.txt")
	defer f.Close()
	if err != nil {
		log.Println("err happened...", err)
	}
}


