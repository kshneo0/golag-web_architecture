package main

import "github.com/web_archi/ch082/fileWriter"

func main() {
	f := fileWriter.NewWriteFile("file.txt")
	f.WriteString("Hello World")
	f.WriteString("More Text!")
	f.Close()

	err := f.Err()
	// var fErr *WriteFileError
	// if errors.As(err, &fErr) {
	// 	if fErr.Op == "Close"{
			
	// 	}
	// }

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