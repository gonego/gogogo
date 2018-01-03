package main
import (
"fmt"
"runtime"
)

func main() {
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Println("To cross compile a Go src file e.g. srcfile.go")
	fmt.Println("Use above info and just type command:")
	fmt.Println("GOOS=...... GOARCH=...... go build srcfile.go")
}
