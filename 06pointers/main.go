package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	// var ptr *int
	// fmt.Println("Pointer value:", ptr) // nil pointer

	myNum := 13
	ptr := &myNum // referencing

	fmt.Println("Pointer value:", ptr)
	fmt.Println("Value at pointer:", *ptr) // dereferencing

}
