package main

import (
	"errors"
	"fmt"
)

func cat() error {
	return errors.New("cat is an error")
}

func moo() error {
	return fmt.Errorf("moo is an error: %w",cat())
}

func bar() error {
	return fmt.Errorf("bar is an error: %w",moo())
}

func foo() error {
	return fmt.Errorf("foo is an error: %w",bar())
	//return fmt.Errorf("foo is an error: %v",bar())
}

func main(){

	// err := foo()
	// fmt.Println(err)

	err := foo()
	baseErr := errors.Unwrap(err)
	fmt.Println(baseErr)

	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)

	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)

	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)
}