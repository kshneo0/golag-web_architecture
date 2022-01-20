package main

import (
	"errors"
	"fmt"
)

func main() {
	fooErr := foo()
	barErr := errors.Unwrap(fooErr)
	mooErr := errors.Unwrap(barErr)
	catErr := errors.Unwrap(mooErr)
	baseErr := errors.Unwrap(catErr)
	fmt.Printf("fooErr\n%s\n", fooErr)
	fmt.Printf("barErr\n%s\n", barErr)
	fmt.Printf("mooErr\n%s\n", mooErr)
	fmt.Printf("baseErr\n%s\n", baseErr)
}

func foo() error {
	return fmt.Errorf("this error is from FOO - %w", bar())
}

func bar() error {
	return fmt.Errorf("this error is from BAR - %w", moo())
}

func moo() error {
	return fmt.Errorf("this error is from MOO - %w", cat())
}

func cat() error {
	return errors.New("this error id from CAT")
}