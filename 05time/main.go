package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time!")

	t := time.Now()

	fmt.Println(t)

	formatedTime := t.Format("01-02-2006 15:04:05 Monday")

	fmt.Println(formatedTime)
}
