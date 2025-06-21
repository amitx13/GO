package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the User Input Program!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Plzz enter ur place to visit:")

	place, _ := reader.ReadString('\n')

	fmt.Println("ur place is:", place)
}
