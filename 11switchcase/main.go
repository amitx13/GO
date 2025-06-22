package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to ludo:")

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	dice := r.Intn(6) + 1

	switch dice {
	case 1:
		fmt.Println("You rolled a 1, move one step forward.")
		fallthrough // fallthrough allows the next case to execute even if it doesn't match
	case 2:
		fmt.Println("You rolled a 2, move two steps forward.")
	case 3:
		fmt.Println("You rolled a 3, move three steps forward.")
	case 4:
		fmt.Println("You rolled a 4, move four steps forward.")
	case 5:
		fmt.Println("You rolled a 5, move five steps forward.")
	case 6:
		fmt.Println("You rolled a 6, move six steps forward and take another turn.")
	default:
		fmt.Println("Invalid roll, please try again.")
	}
}
