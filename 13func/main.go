package main

import "fmt"

func main() {
	var x interface{} = "apx"
	str, ok := x.(string)
	if ok {
		fmt.Println("utsav ki taiyari kro string aa rahe h:", str)
	} else {
		fmt.Println(x)
	}
}

//what is type assertion in go?
// Type assertion is a way to retrieve the dynamic type of an interface in Go.
// It allows you to check if the interface holds a specific type and, if so, extract that value.
// The syntax for type assertion is `value, ok := interfaceValue.(Type)`, where `Type` is the type you want to assert.
// If the assertion is successful, `ok` will be true, and `value` will hold the asserted value.
// If the assertion fails, `ok` will be false, and `value` will be the zero value of the asserted type.
