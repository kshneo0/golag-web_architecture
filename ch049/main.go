package main

import "fmt"

type hotdog int

func (f hotdog) cook() {
	fmt.Println("I am cooking")
}

type hotFood interface {
	cook()
}

func bar(hf hotFood) {
	hf.cook()
}

func main() {
	var x hotdog = 42
	var y hotFood
	y = x

	fmt.Printf("%T\n",y)
	z := hotFood(x)
	fmt.Printf("%T\n",z)

	bar(x)
}