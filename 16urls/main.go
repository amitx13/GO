package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://example.com/learn?topic=golang&level=beginner"

func main() {
	fmt.Println("Welcome to Urls in Go!")

	res, _ := url.Parse(myUrl)

	fmt.Println("Scheme", res.Scheme)
	fmt.Println("Host", res.Host)
	fmt.Println("port", res.Port())
	fmt.Println("Path", res.Path)
	fmt.Println("RawQuery", res.RawQuery)

	querys := res.Query()

	for i, val := range querys {
		fmt.Printf("Key is %v and val is %v\n", i, val)
	}

	reconstructUrl := &url.URL{
		Scheme:   res.Scheme,
		Host:     res.Host,
		Path:     res.Path,
		RawQuery: res.RawQuery,
	}
	fmt.Println("Reconstructed URL:", reconstructUrl.String())
}
