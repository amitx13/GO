package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Dominos Pizza!")

	fmt.Println("Please input your rating (1-5):")

	input := bufio.NewReader(os.Stdin)

	rating, _ := input.ReadString('\n')

	fmt.Println("Your rating is:", rating)

	//Hack add one extra to user rating

	updatedRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println("Error converting rating:", err)
		return
	}

	fmt.Println("Your updated rating is:", updatedRating+1)

}
