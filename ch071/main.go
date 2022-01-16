package main

import (
	"errors"
	"fmt"
)

var ErrNotExist = fmt.Errorf("File does not exist")
var ErrUserNotExist = errors.New("User does not exist")

func main() {
	// var err = fmt.Errorf("File not found: %s", "filename")

	err := ErrUserNotExist

	if err == ErrUserNotExist {
		fmt.Println("You nedd to register first!")
	} else {
		fmt.Println("Unknown error")
	}
}