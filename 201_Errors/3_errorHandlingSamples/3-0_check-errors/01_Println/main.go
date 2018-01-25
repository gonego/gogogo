package main

import "fmt"

func main() {
	name, age, country := info()
	n, err := fmt.Println("Hello", name, ",", age, "from", country)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("No of bytes:", n)

}
func info() (string, string, string) {
	var name, age, country string

	fmt.Print("Name: ")
	//number of items scan is discarded
	_, err := fmt.Scan(&name)
	if err != nil {
		panic(err)
	}

	fmt.Print("Age: ")
	_, err = fmt.Scan(&age)
	if err != nil {
		panic(err)
	}

	fmt.Print("Country: ")
	_, err = fmt.Scan(&country)
	if err != nil {
		panic(err)
	}

	fmt.Println("Age:", age, "From:", country)
	return name, age, country
}
