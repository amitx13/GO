package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("Welcome to web Requests in Go")

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(errors.New("error while making the request"))
		return
	}

	fmt.Println("Body", res.Body)

	constents, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(errors.New("error while reading the response body"))
		return
	}

	fmt.Println("res", string(constents))
	defer res.Body.Close()
}
