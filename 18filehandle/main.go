package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handle Secition of this go series")

	// We os and io things

	content := "This is the string to be inside the file"
	//Create a file First using os module
	file, err := os.Create("./newFile.txt")

	checkError(err)

	//To write
	length, err := io.WriteString(file, content)
	checkError(err)
	defer file.Close()
	fmt.Println("Length of the file is:", length)
	readfile("./newFile.txt")

}

func readfile(filename string) {

	dataBytes, err := os.ReadFile(filename)
	checkError(err)
	fmt.Println("Length of the file is:", string(dataBytes))

}

func checkError(err error) {

	if err != nil {
		panic(err)

	}

}
