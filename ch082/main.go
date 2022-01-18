package main

import (
	"fmt"
	"io"
	"os"
)
type writeFileError struct{
	Op string
	Err error
}

func (w writeFileError) Error() string {
	return w.Err.Error()
}

func (w writeFileError) Unwrap() error {
	return w.Err
}

type writeFile struct {
	f * os.File
	err error
}

func newWriteFile(filename string) *writeFile {
	f, err := os.Create(filename)
	return &writeFile{
		f: f,
		err: err,
	}
}

func (w *writeFile) WriteString(text string) {
	if w.err != nil{
		return
	}

	_, err := io.WriteString(w.f, text)
	if err != nil {
		//w.err = err
		w.err = writeFileError {
			Op: "WritieString",
			Err: fmt.Errorf("Failed wile writing a string: %w", err),
		}
		//w.err = fmt.Errorf("Failed wile writing a string: %w", err)
	}
}

func (w *writeFile) Close() {
	if w.err != nil {
		return
	}

	err := w.f.Close()
	if err != nil {
		w.err = err
	}
}

// All errors returning from Err should be of type *WriteFileError
func (w *writeFile) Err() error {
	return w.err
}

func main() {
	f := newWriteFile("file.txt")
	f.WriteString("Hello World")
	f.WriteString("More Text!")
	f.Close()

	err := f.Err()
	if err != nil {
		panic(err)
	}

	// f, err := os.Create("file.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = io.WriteString(f, "Hello world")
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = io.WriteString(f, "More Text!")
	// if err != nil {
	// 	panic(err)
	// }
	// err = f.Close()
	// if err != nil {
	// 	panic(err)
	// }
}