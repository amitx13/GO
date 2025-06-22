package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Go")

	lang := make(map[string]string)

	lang["RS"] = "Rust"
	lang["GO"] = "GoLang"
	lang["JS"] = "JavaScript"

	fmt.Println("Languages:", lang)
	fmt.Println("RS stands for:", lang["RS"])

	delete(lang, "JS")
	fmt.Println("Languages:", lang)

	//foreach loop

	for key, val := range lang {
		fmt.Printf("For key %v, value is %v\n", key, val)
	}
}
