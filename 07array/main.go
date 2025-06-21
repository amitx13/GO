package main

import "fmt"

func main() {
	var arr [4]string

	arr[0] = "Hello"
	arr[1] = "World"
	arr[3] = "Go"

	fmt.Println("Array values:", arr)
	fmt.Println("Array length:", len(arr))

	var arr2 = [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Array 2 values:", arr2)
	fmt.Println("Array 2 length:", len(arr2))
}
