package _18_json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	//func Open(name string) (*File, error)
	//*File is attached with methods:
	//func (f *File) Write(b []byte) (n int, err error)
	//func (f *File) Read(b []byte) (n int, err error)
	//to implement both io.Reader and io.Writer
	fileReader, _ := os.Open("vee.json")

	data := make([]byte, 200)
	count, err := fileReader.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes:\n%s\n", count, data[:count])

	fileReader2, _ := os.Open("vee.json")
	if err != nil {
		log.Fatal(err)
	}
	var v map[string]interface{}
	json.NewDecoder(fileReader2).Decode(&v)
	fmt.Println(v)

}
