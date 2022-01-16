package main

import (
	"errors"
	"fmt"
	"time"
)

type ErrFileNotFound struct {
	Filename string
	When time.Time
}

func ( e ErrFileNotFound) Error() string {
	return fmt.Sprintf("File %s was not found at %v", e.Filename, e.When)
}

func (w ErrFileNotFound) Is(other error) bool {
	_, ok := other.(ErrFileNotFound)
	return ok
}

var ErrNotExist = fmt.Errorf("File does not exist")
var ErrUserNotExist = errors.New("User does not exist")

func openFile(filename string) (string, error) {
	return "", ErrNotExist
}

func main() {
	_, err := openFile("test.txt")
	if err != nil {
		wrappedErr := fmt.Errorf("Unable to open file %v: %w","test.txt",err)
		if errors.Is(wrappedErr, ErrNotExist){
			fmt.Println("This is an ErrNotExist")
		}
		panic(wrappedErr)
	}
}