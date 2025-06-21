package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices in Go")

	var sli = []string{"apple", "banana", "cherry"}
	fmt.Println(sli)

	sli = append(sli, "peach", "grape")
	fmt.Println(sli)

	// list := make([]int, 4)

	// list[0] = 97
	// list[1] = 62
	// list[2] = 73
	// list[3] = 44

	// fmt.Println(list)

	// list = append(list, 99, 100) // Append more elements automatically resizes the slice
	// fmt.Println(list)

	// sort.Ints(list)

	// fmt.Println("Sorted list:", list)

	//how to remove a value from slices based on index
	//remove index 2

	indexToRemove := 2

	sli = append(sli[:indexToRemove], sli[indexToRemove+1:]...)

	fmt.Println("After removing index 2:", sli)
}
