package main

import "fmt"

type person struct {
	first string
}

func (p person) String() string {
	return fmt.Sprint("My name is ", p.first)
}

func main() {
	p1 := person{
		first: "Jenny",
	}
	fmt.Println(p1)

	// go's example
	a := Animal{
		Name: "Gopher",
		Age: 2,
	}

	fmt.Println(a)
}

// Animal has a Name ans an Age ri represent an animal.
type Animal struct {
	Name string
	Age uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}