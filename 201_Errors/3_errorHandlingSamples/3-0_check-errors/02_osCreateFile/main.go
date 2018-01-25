package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("wta2018.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader(`
	Eugenia Bouchard seems to have lost the speed in her court coverage.
	Her lateral movement from side to side is somewhat sluggish and is
	starkly lagging behind her shot making quickness.`)

	n, err := io.Copy(f, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Bytes written:", n)

}
