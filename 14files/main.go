package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Go!")

	contents := "This need to go inside a fucking file."

	file, err := os.Create("./myGoFile.txt")

	if err != nil {
		fmt.Println(errors.New("Error creating file: " + err.Error()))
	}

	len, err := io.WriteString(file, contents)

	if err != nil {
		fmt.Println(errors.New("Error writing to file: " + err.Error()))
	}

	fmt.Printf("Length of file is %d bytes\n", len)

	defer file.Close()

	readFile("myGoFile.txt")
}

func readFile(fileName string) {

	databyte, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(errors.New("Error reading file: " + err.Error()))
		return
	}
	fmt.Println("File content is: ", string(databyte))

}
