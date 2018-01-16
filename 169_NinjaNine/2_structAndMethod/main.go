package main

import "fmt"

type Person struct {
	Name        string
	Age         int
	Nationality string
}

//a method with pointer type receiver
//can be accessed both by pointer and
//non-pointer type
//however if this method is used to
//implement interface type, only a pointer
//can access any function with the interface
//type parameter
func (p *Person) Info() {
	fmt.Println(*p)
}

type Human interface {
	Info()
}

//note that an interface type is not
//accepted as receiver type. Thus
//we use func instead of method here
func MoreInfo(h Human) {
	fmt.Println()
	h.Info()
	switch h.(*Person).Name {
	case "Linus Torvalds":
		fmt.Println("Created Linux and Git")
	case "Rob Pike":
		fmt.Println("Created Golang & UTF8")
	case "Ken Thompson":
		fmt.Println("Created Unix, UTF8 & Golang")
	default:
		fmt.Println("Must be somebody important")
	}
	fmt.Println()
}

func AllInfo(i ...interface{}) {
	fmt.Println(i)
}

func main() {
	p1 := Person{
		"Linus Torvalds",
		48,
		"Finland",
	}

	p2 := Person{
		"Rob Pike",
		61,
		"Canada",
	}

	p3 := Person{
		"Ken Thompson",
		74,
		"USA",
	}

	//non-pointers can invoke Info()
	//even though it's a method with
	//pointer receiver
	p1.Info()
	p2.Info()
	p3.Info()

	//pp1 a pointer to p1
	pp1 := &p1

	//a pointer can of course invoke Info()
	//which accepts pointer receiver
	pp1.Info()

	//However only a pointer can be passed as
	//argument to func MoreInfo()which only accept
	//interface type. This is because the
	//method to qualify Person as an interface
	//to Human, Info() has a pointer type receiver
	//MoreInfo(p1) is not acceptable
	MoreInfo(pp1)

	//AllInfo(i ...interface{})
	//variadic parametere of any type
	//These will create []interface{}
	AllInfo(&p1, &p2, p3)
	AllInfo(pp1, pp1, p3)

}
