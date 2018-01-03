package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	p := Person{
		Name:   "Fiona Gallagher",
		Age:    27,
		Emails: []string{"fna@gmail.com", "fna@yahoo.com"},
	}
	//func Create(name string) (*File, error)
	fileWriter, _ := os.Create("fiona.json")
	//func NewEncoder(w io.Writer) *Encoder
	//func (enc *Encoder) Encode(v interface{}) error
	json.NewEncoder(fileWriter).Encode(p)
}