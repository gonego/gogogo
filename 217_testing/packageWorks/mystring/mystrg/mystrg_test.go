package mystrg

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "Go is all about type"
	x := strings.Split(s, " ")
	s = Cat(x)
	if s != "Go is all about type" {
		t.Error("got", s, "want", "Go is all about type")
	}
}

func TestJoin(t *testing.T) {
	s := "Go is all about type"
	x := strings.Split(s, " ")
	s = Join(x)
	if s != "Go is all about type" {
		t.Error("got", s, "want", "Go is all about type")
	}
}

func ExampleCat() {
	s := "Go is all about type"
	x := strings.Split(s, " ")
	fmt.Println(Cat(x))
	// Output:
	// Go is all about type
}

func ExampleJoin() {
	s := "Go is all about type"
	x := strings.Split(s, " ")
	fmt.Println(Join(x))
	// Output:
	// Go is all about type
}

const s = "Living in the fast or slow lane is completely up to you. " +
	"Only understand that life has many detour and no turning back. " +
	"Reaching your destination too soon might lose you the enjoyment" +
	"of having a good view of the journey."

var x []string

func BenchmarkCat(b *testing.B) {
	x = strings.Split(s, " ")
	//to exclude time for Split
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(x)
	}
}

func BenchmarkJoin(b *testing.B) {
	x = strings.Split(s, " ")
	//to exclude time for Split
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(x)
	}
}
