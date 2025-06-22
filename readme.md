# GOOOOOOOO

# WHY GO?
```go
1. Go is a statically typed, compiled language. which means that the type of a variable is known at compile time.
2. Fast and lightweight
3. Easily concurrent(Multiple tasks make progress at the same time, but they may not run simultaneously. Think of it as task switching) and (parallel)multi-threaded
4. Garbage collection
5. Go code generally runs faster than interpreted languages (Js, python, Ruby) and compiles faster than other compiled languages (Rust, C, C++, etc.)
6. Go runtime is responsible for managing memory (to cleanup unused memory), garbage collection, and scheduling goroutines.
```
# Lexer
The Job of the lexer is to simple understand that u are following the grammer of the lang. So that the syntax is correct it does it before compilation 

# Memory Management 
Memory allocation and deallocation happens automatically in GO.
*new()* - Allocate memory but no initialization, will get a memory address, zeroed storage - we cannot put any data initialy
*make()* - Allocate memory and init, will get a memory address, non-zeroed storage - can put any data.
*GC* - Garbage collector happens automatically anything that is nil or outof scope is eligible.

# GO ENV (go env)
*GOGC* - The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly collected data to live data remaining after the previous collection reaches this percentage. The dafault is GOGC=100

-: *Notes* :-

Case insensitive.
Variable type should be known in advance.
Almost Everything is a Type in Go.
Formating string in Go:
    %v - represents the value use as a default format
    %d - represents an integer
    %f - represents a floating-point number
    %t - represents a boolean
    %s - represents a string
    %q - represents a string as it is quoted
    %T - represents the type of the value

# Conditionals

```go
if size > 5 {
    fmt.Println("You are large enough!")
}

we can also define variable in the initial if stmt
if size := 7; size > 5{
    fmt.Println("f u are big")
}
```

why pointer?

Sometimes when u pass on these variable a copy of this variable is passed on whenever there is a case when u want to avoide such this cases and u want absolute
garentee that always actual value should be passed on then we prefer that a pointer should be passed. A pointer is just a direct reference to a memory address

# Switch case

Switch statements are a way to compare a value against multiple options

```go
func getCreator(name string) string {
	switch name {
	case "linux":
		return "Linus Torvalds"
	case "mac":
		return "Steve Jobs"
	case "SpaceX":
		fallthrough //fallthrough allows the next case to execute even if it doesn't match
	case "Tesla":
		return "Elon Musk"

	default:
		return "Unknown"
	}
}
```

# Functions

functions in Go is defined using `func` keyword 
*Types of return in Go*
1. Named return 
2. Naked return


# Named Return Example

```go
func add(a int, b int) (sum int) {
    sum = a + b
    return // returns the named variable 'sum'
}
```

# Naked Return Example

```go
func multiply(a int, b int) (int) {
    return a * b // directly returns the result
}
```

*Guard clauses are a way to handle errors in Go. They are a way to handle errors early in a function and return early if an error occurs.*

# Anonymous function 
Anonymous functions are functions that are defined without a name. They are useful when you need to define a function inline without having to name it.
Or can only be used once.

```go
func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

func double(a int) int {
    return a + a
}

func main() {
	newX, newY, newZ := conversions(double, 1, 2, 3)
	// newX is 2, newY is 4, newZ is 6

    //anonymous func
	newX, newY, newZ = conversions(func(a int) int {
	    return a * a
	}, 1, 2, 3)
	// newX is 1, newY is 4, newZ is 9
}
```

# Defer
The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns.
Defer is often used to ensure that resources are released when they are no longer needed.
Defer is a great way to make sure that something happens before a function exits, even if there are multiple return statements, a common occurrence in Go.
a func can have mutiple defer statements and they are executed in LIFO order.

usecase: defer fmt.Println("TEXTIO BOOTUP DONE")

# Closures
Closures in Go is similar to JavaScript. A closure is a function that has access to variables from an outer function in which it was declared.
It can access variables from the outer function even after the outer function has finished executing.

```go
func main(){
		counter := 0
		return func() int { // anonymous Closures function
			counter ++
			return counter
		}
}
```

# Methods

*Methods are just fuction that have a receiver. A receiver helps buinding a specific method to a specific type.*

There are two types of receiver in Go:
1. Value receiver
2. Pointer receiver

```go
package main

import "fmt"

	type rect struct {
		height, weigth int
	}

func (r rect) area() int { // value receiver

		return r.height * r.weigth
	}

func (r *rect) scale(parameter int) { // pointer receiver

		r.height *= parameter
		r.weigth *= parameter
	}

	func main() {
		r := rect{height: 10, weigth: 5}
		fmt.Println("Area before scale:", r.area())
		r.scale(2)
		fmt.Println("Area after scale:", r.area())
	}

package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) varify() string { //valur reciver
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func main() {
	auth := authenticationInfo{
		username: "xlr8",
		password: "1234",
	}
	fmt.Println(auth.varify())
}

```

# Memory Layout

*Why the ordering of fields in a struct matters?*
becz it can have huge impact on memory usage and performance.
Always try to order fields in a struct from largest to smallest to minimize memory usage.

```go 
int > float64 > string > bool
```

# Interfaces

Interface is a type the only specify the set of methods but does not provide it's implementation.
It's just collection of methods signature.

Note :- 
`A type that implements all the methods of an interface implements interface.`
`Interfaces are implemented implicitly.`
`A quick way of checking whether a struct implements an interface is to declare a function that takes an interface as an argument. If the function can take the struct as an argument, then the struct implements the interface.`
`Keep interface small.`

# Additional informations about interfaces:

1. The empty interface - interface{} It does not have any methods. Therefore, every types in GO implements empty interface.
2. Zero value of an interface - It can be nil. It's their zero value. That's why when we check for error in Go, we check err != nil
becz err is an interface. It represents error that doesn't exist.
3. Interfaces on pointer -

example:
```go
this two are two different things

type rectangle interface {
    height() int
    width() int
}
type square struct {
    length int
}
func (sq *square) width() int {
    return sq.length
}
func (sq *square) height() int {
    return sq.length
}

Here square does not implemented the rectangle interface the *square does.
if it was
func (sq square) width() int {
    return sq.length
}

func (sq square) height() int {
    return sq.length
}

then square would have implemented the rectangle interface.

package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func main() {
	emp := contractor{
		name:         "apx13",
		hourlyPay:    10,
		hoursPerYear: 1440,
	}
	fmt.Println("Name of the employee:", emp.getName())
	fmt.Println("Yearly package:", emp.getSalary())
}
```

# Type Interface in GO:
An interface itself does not store data directly, but an interface variable can hold a value and it's type
Type assertions help you "extract" that value when needed.

```go
package main

import "fmt"

func main() {
	var x interface{} = 42
	str, ok := x.(string)
	if ok {
		fmt.Println("utsav ki taiyari kro string aa rahe h:", str)
	} else {
		fmt.Println(x)
	}
}
```

```go
Type switches - A type switch makes it easy to do several type assertions in a series.

package main
import "fmt"

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is int:", v)
	case string:
		fmt.Println("IT's a string:", v)
	case float64:
		fmt.Println("F it's a float:", v)
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	printType(42)
	printType("Go")
	printType(3.14)
}
```

The Error Interface:

1. When something can go wrong in a function, that function should return an error as its last return value.
2. Any code that calls a function that can return an error should handle errors by testing whether the error is nil
3. A nil error denotes success; a non-nil error denotes failure.
4. Note: When you return a non-nil error in Go, it's conventional to return the "zero" values of all other return values.
5. Genral rule of thumb: Do not use Panic. becz When a function calls panic, the program crashes and prints a stack trace.

The Go standard library provides an "error" package
/*
Use case:
	import "errors"

	errors.new("Something Went wrong!") // Don't forget to return 0 in case of error.
*/
`If you want your program to cleanly exit in an unrecoverable way, which is a good alternative to panic? log.Fatal()`

```go
package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func validateStatus(status string) error {
	len := utf8.RuneCountInString(status)
	if len == 0 {
		return errors.New("status cannot be empty")
	}

	if len > 140 {
		return errors.New("status exceeds 140 characters")
	}

	return nil
}

func main() {
	res := validateStatus(":)")
	if res != nil {
		fmt.Println("Something went wrong! Error:", res)
		return
	}
	fmt.Println("Status Updated")
}
```
