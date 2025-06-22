package main

import "fmt"

func main() {

	apx := User{"apx", "apxkvs981@gmail.com", 25, true}

	fmt.Printf("User details: %+v\n", apx)
}

type User struct {
	Name       string
	Email      string
	Age        int
	IsVerified bool
}
