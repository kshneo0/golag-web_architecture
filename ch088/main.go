package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	src := "kabir.txt"
	dst := "second-file.txt"
	err := copyFile(dst, src)
	if errors.Is(err, os.ErrNotExist){
		fmt.Println("You need to provide the name \"kabir.txt\" of a valid file in your directory next to the executable")
	} else if err != nil {
		log.Panicln("in main, caling copyFile return an error:", err)
	}
}

func copyFile(dst, src string) error {
	f1, err:= os.Open(src)
	if err != nil{
		return fmt.Errorf("couldn't open file in CopyFile: %w", err)
	}
	defer f1.Close()

	f2, err := os.Create(dst)
	// log.Panic("for the fun of it")
	if err != nil {
		return fmt.Errorf("couldn't create a dile in CopyFile: %w", err)
	}
	defer f2.Close()

	n, err := io.Copy(f2, f1)
	if err != nil {
		return fmt.Errorf("couldn't copy a file in CopyFile: %w", err)
	}
	fmt.Println("just in development, nice to see bytes written:", n)

	return nil
}