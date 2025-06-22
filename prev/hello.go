// `go run fileName` to compiled and run
// `go build fileName` to build and then exec the file directly

// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	a := "xlr8"
// 	fmt.Println("Hello, World!", a)
// 	str := utf8.RuneCountInString(a)
// 	fmt.Printf("The length of %s is %d\n", a, str)
// }

// walrus operator `:=` is used to declare and initialize a variable at the same time. The limitation is that := can’t be used outside of a function
// := requires at least one new variable to be declared. Cannot be used with Predefined Types (struct, fields) also cannot be used In Function Signature or Method Receivers

//constants const a = 10

//Why Go?
// 1. Go is a statically typed, compiled language. which means that the type of a variable is known at compile time.
// Fast and lightweight
// Easily concurrent(Multiple tasks make progress at the same time, but they may not run simultaneously. Think of it as task switching) and (parallel)multi-threaded
// Garbage collection
// Go code generally runs faster than interpreted languages (Js, python, Ruby) and compiles faster than other compiled languages (Rust, C, C++, etc.)

/*
Default types:

bool
string
int
uint
byte
rune
float64
complex128

Note: The only place you do not want to use default types is when you want to squeeze out every last bit of performance and memory usage. And there is a specific need to do that.
Most of the time developers use default types because they are easier to read and understand.
*/

// Benefits of using compiled language:
/*
You can run a compiled program without the original source code.
You don’t need the compiler anymore after it’s done its job.
That’s how most video games are distributed! Players don’t need to install the correct version of Go to run a PC game:
they just download the executable game and run it.
*/

// Go runtime is responsible for managing memory (to cleanup unused memory), garbage collection, and scheduling goroutines.

/*
Constants:
Constants are declared with the const keyword. They can’t use the := (walrus operator) short declaration syntax.
Constants can be primitive types like strings, integers, booleans and floats. They can not be more complex types like slices, maps and structs,
The reason why const in go not allowed to use walrus := is becz in go const represent a constant value not a variable.
However, constants can be computed as long as the computation can happen at compile time.

That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:
// the current time can only be known when the program is running
const currentTime = time.Now()

*/

/*
Formating string in Go:
%v - represents the value use as a default format
%d - represents an integer
%f - represents a floating-point number
%t - represents a boolean
%s - represents a string
%q - represents a string as it is quoted
%T - represents the type of the value
*/
// When you need to work with individual characters in a string, you should use the rune type(utf8.RuneCountInString("apx13")). It breaks strings up into their individual characters, which can be more than one byte long.

/*
Conditionals:

if statements in Go do not use parentheses around the condition:
if height > 4 {
    fmt.Println("You are tall enough!")
}

Initial smt in if statement:
we can define initial variable in an if smt and the variable created in the initial smt are only defined within the scope of the if body.

Like instead of writing:
	size := 7
	if size > 8 {
		fmt.Println("Big")
	}

Write:
if size := 7; size > 8 {
	fmt.Println("Big")
}
	1. it's a bit shorter
	2. the size variable is only defined within the scope of the if body

Switch statement:
Switch statements are a way to compare a value against multiple options

func getCreator(name string) string {
	switch name {
	case "linux":
		return "Linus Torvalds"
	case "mac":
		return "Steve Jobs"
	case "SpaceX":
		fallthrough
	case "Tesla":
		return "Elon Musk"

	default:
		return "Unknown"
	}
}

Note: the breack statement is not required in Go. Go automatically (implecit) breaks out of the switch statement after a case is matched.
fallthrough allows the next case to execute even if it doesn't match
*/

/*
Functions:

Code example:
func sub(x int, y int) int {
  return x-y
}
Type of Return in Go:
 Named return
 Naked return

 Note: A return statement without arguments returns the named return values. This is known as a “naked” return.
 Naked return statements should be used only in short functions. They can harm readability in longer functions.

func getCoords() (x, y int){
  // x and y are initialized with zero values

  return // automatically returns x and y
}

Is the same as:

func getCoords() (int, int){
  var x int
  var y int
  return x, y
}

*/

/*
Random 3am Go stuff
Note: nil is the zero value of an error interface. It’s a special value that represents an error that doesn’t exist.
*/

/*
Guard Clauses:
Guard clauses are a way to handle errors in Go. They are a way to handle errors early in a function and return early if an error occurs.

Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional.
Instead of using if/else chains, we just return early from the function at the end of each conditional block.

Example
package main

import "fmt"

	func main() {
		divideBoth(10, 0)
	}

	func divideBoth(x, y int) int {
		if y == 0 { // Guard clauses: Handle error cases early.
		fmt.Println("Cannot divide by 0")
		return 0
		} else { //Main logic comes after handling edge cases.
		fmt.Println("output:", x/y)
		return x / y
	}

}

package main

import (
	"errors"
	"fmt"
)

func validateUser(age int, name string) error {
	if age < 18 {
		return errors.New("user is under 18")
	}

	if name == "" {
		return errors.New("name is empty")
	}

	fmt.Println("User is valid!")
	return nil
}

func main() {
	if err := validateUser(16, "John"); err != nil {
		fmt.Println("Validation failed:", err)
	}

	if err := validateUser(25, ""); err != nil {
		fmt.Println("Validation failed:", err)
	}

	validateUser(25, "Alice") // Output: User is valid!
}

*/

/*
Anonymous functions:
Anonymous functions are functions that are defined without a name. They are useful when you need to define a function inline without having to name it.
Or can only be used once.

example:
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
    // using a named function
	newX, newY, newZ := conversions(double, 1, 2, 3)
	// newX is 2, newY is 4, newZ is 6

    // using an anonymous function
	newX, newY, newZ = conversions(func(a int) int {
	    return a * a
	}, 1, 2, 3)
	// newX is 1, newY is 4, newZ is 9
}
*/
/*
Defer
The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns.
Defer is often used to ensure that resources are released when they are no longer needed.
Defer is a great way to make sure that something happens before a function exits, even if there are multiple return statements, a common occurrence in Go.

usecase: defer fmt.Println("TEXTIO BOOTUP DONE")

a func can have mutiple defer statements and they are executed in LIFO order.


Closures in Go is similar to JavaScript. A closure is a function that has access to variables from an outer function in which it was declared.
It can access variables from the outer function even after the outer function has finished executing.

example:

func main(){
		counter := 0
		return func() int { // anonymous Closures function
			counter ++
			return counter
		}
}
*/

/*
Structs:
package main

	type car struct {
		brand string
		model string
		milage int
	}

Nested structs:
package main

import "fmt"

	type messageToSend struct {
		message  string
		sender   user
		recipent user
	}

	type user struct {
		name   string
		number int
	}

	func canSendMessage(mToSend messageToSend) bool {
		if mToSend.message == "" {
			return false
		}

		if mToSend.sender.name == "" || mToSend.recipent.name == "" {
			return false
		}

		if mToSend.sender.number == 0 || mToSend.recipent.number == 0 {
			return false
		}

		return true
	}

	func main() {
		mToSend := messageToSend{
			message:  "you have an appointment tomorrow",
			sender:   user{name: "Brenda Halafax", number: 16545550987},
			recipent: user{name: "Sally Sue", number: 19035558973},
		}

		mNotToSend := messageToSend{
			message:  "",
			sender:   user{name: "", number: 16545550987},
			recipent: user{name: "Sally Sue", number: 0},
		}

		fmt.Println("Message status:", canSendMessage(mToSend))
		fmt.Println("Message status:", canSendMessage(mNotToSend))

}

Anonymous structs:

a anonymous struct is just like a normal struct but without a name. since it does not have a name so it cannot be referenced later on.
Hence the anonymous struct can only be initialize at time of decleration.

use case: only use it in case you are sure that u are not going to reuse it again

package main

import "fmt"

	func main() {
		car := struct {
			brand string
			model string
		}{
			brand: "Dogde",
			model: "Charger",
		}

		fmt.Println(car)
	}

You can even nest anonymous structs as fields within other structs:

	type car struct {
	  brand string
	  model string
	  doors int
	  mileage int
	  // wheel is a field containing an anonymous struct
	  wheel struct {
	    radius int
	    material string
	  }
	}

Embedded structs:

package main

import "fmt"

	type car struct {
		brand string
		model string
	}

	type truck struct {
		car
		bedSize int
	}

	func main() {
		t := truck{
			car: car{
				brand: "Ford",
				model: "F150",
			},
			bedSize: 8,
		}

		fmt.Println(t.brand, t.model, t.bedSize)

		// if it was nested then i have to go more deeper to access the fields
		// fmt.Println(t.car.brand, t.car.model, t.bedSize)
		// but with embedded struct i can access the fields directly
	}

Methods:

Methods are just fuction that have a receiver. A receiver helps buinding a specific method to a specific type.

There are two types of receiver in Go:
1. Value receiver
2. Pointer receiver

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

func (a authenticationInfo) varify() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func main() {
	auth := authenticationInfo{
		username: "xlr8",
		password: "1234",
	}
	fmt.Println(auth.varify())
}

Memory Laypout: (Imp for performance)
the memory layout in Go is similar to C.
In Go, structs sit in memory in a contiguous block, with fields placed one after another as defined in the struct.

type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

Reach | NumPosts | NumLikes

Why the ordering of fields in a struct matters?
becz it can have huge impact on memory usage and performance.
Always try to order fields in a struct from largest to smallest to minimize memory usage.

int > float64 > string > bool

package main

import "fmt"

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	name string
	Membership
}

func newUser(name string, membershipType string) User {
	if membershipType == "premium" {
		return User{
			name: name,
			Membership: Membership{
				Type:             membershipType,
				MessageCharLimit: 1000,
			},
		}
	} else {
		return User{
			name: name,
			Membership: Membership{
				Type:             membershipType,
				MessageCharLimit: 100,
			},
		}
	}
}

func main() {
	user1 := newUser("apx13", "premium")
	user2 := newUser("xlr8", "basic")

	fmt.Println("Premium User:", user1)
	fmt.Println("Basic User", user2)
}
*/

/*
# Interfaces in Go:

Interface is a type that specify a set of methods but does not provide it's implementation. It allow you to define behaviour without worring about how's it done
It's just collection of methods signature.
A type that implements all the methods of an interface implements interface
Means - Interfaces are implemented implicitly.

A quick way of checking whether a struct implements an interface is to declare a function that takes an interface as an argument.
If the function can take the struct as an argument, then the struct implements the interface.

# Thumb rules:
1. Keep Interface Small.
2. Interfaces Should Have No Knowledge of Satisfying Types.
	An interface should define what is necessary for other types to classify as a member of that interface.
	They shouldn’t be aware of any types that happen to satisfy the interface at design time.
3. Interfaces are not classes
		- Interfaces aren’t hierarchical by nature, though there is syntactic sugar to create interfaces that happen to be supersets of other interfaces.
		- Interfaces define function signatures, but not underlying behavior.


# Additional informations about interfaces:

1. The empty interface - interface{} It does not have any methods. Therefore, every types in GO implements empty interface.
2. Zero value of an interface - It can be nil. It's their zero value. That's why when we check for error in Go, we check err != nil
becz err is an interface. It represents error that doesn't exist.
3. Interfaces on pointer -

example:

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

*/
/* package main

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
*/
/*
Type Interface in GO:
An interface itself does not store data directly, but an interface variable can hold a value and it's type
Type assertions help you "extract" that value when needed.
*/

//Type Assertion
// package main

// import "fmt"

// func main() {
// 	var x interface{} = 42
// 	str, ok := x.(string)
// 	if ok {
// 		fmt.Println("utsav ki taiyari kro string aa rahe h:", str)
// 	} else {
// 		fmt.Println(x)
// 	}
// }

// Type switches - A type switch makes it easy to do several type assertions in a series.
// package main

// import "fmt"

// func printType(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println("This is int:", v)
// 	case string:
// 		fmt.Println("IT's a string:", v)
// 	case float64:
// 		fmt.Println("F it's a float:", v)
// 	default:
// 		fmt.Println("unknown type")
// 	}
// }

// func main() {
// 	printType(42)
// 	printType("Go")
// 	printType(3.14)
// }

//The Error Interface:

// When something can go wrong in a function, that function should return an error as its last return value.
// Any code that calls a function that can return an error should handle errors by testing whether the error is nil
// A nil error denotes success; a non-nil error denotes failure.
// Note: When you return a non-nil error in Go, it's conventional to return the "zero" values of all other return values.
// Genral rule of thumb: Do not use Panic. becz When a function calls panic, the program crashes and prints a stack trace.

// The Go standard library provides an "error" package
/*
Use case:
	import "errors"

	errors.new("Something Went wrong!") // Don't forget to return 0 in case of error.
*/
// If you want your program to cleanly exit in an unrecoverable way, which is a good alternative to panic? log.Fatal()

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"unicode/utf8"
// )

// func validateStatus(status string) error {
// 	len := utf8.RuneCountInString(status)
// 	if len == 0 {
// 		return errors.New("status cannot be empty")
// 	}

// 	if len > 140 {
// 		return errors.New("status exceeds 140 characters")
// 	}

// 	return nil
// }

// func main() {
// 	res := validateStatus(":)")
// 	if res != nil {
// 		fmt.Println("Something went wrong! Error:", res)
// 		return
// 	}
// 	fmt.Println("Status Updated")
// }

/*
Loops:

for INITIAL ; CONDITION ; AFTER {
	// do something
}

1 . loops in go can omit sections of a for loop
ex:

for init ; ; af {
		// infinit loop
}

2 . There is no while loop in GO but since we can omit sections of a for loop we can create a while from a for loop
ex: A while loop is just a for loop what only has condition.

for CONDITIONS {
  // do some stuff while CONDITION is true
}

*/

/* package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
*/

/*
Slices:

Array in go is fixed size groups of variable of same type.
Slices are much better than array in go it's dynamic
Slice does not store it's own data but instead refrence an underlying array
most of the time we will use slices when working with ordered lists.

Note: Slices and map are passed by referance into function

slices ex:
arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]
Copy icon
Where lowIndex is inclusive and highIndex is exclusive.

lowIndex, highIndex, or both can be omitted to use the entire array on that side of the colon.

len and cap are legal when applied to the nil slice, and return 0.
*/
// package main

// import "fmt"

// func main() {
// 	//init array
// 	var arr [5]int

// 	//init slice
// 	var sli []int
// 	fmt.Println(arr)
// 	fmt.Println(sli)

// 	//dec & init arr and slice
// 	arr2 := [3]int{1, 2, 3}
// 	sli2 := []int{1, 2, 3}

// 	fmt.Println(arr2)
// 	fmt.Println(sli2)

// 	//using make
// 	sli3 := make([]int, 3) //init a slice of len 3 with every value to 0
// 	fmt.Println(sli3)
// }

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slice := arr[1:4] // Slice from arr[1] to arr[3]

// 	fmt.Println("Slice:", slice)         // [2 3 4]
// 	fmt.Println("Length:", len(slice))   // 3 (2, 3, 4)
// 	fmt.Println("Capacity:", cap(slice)) // 4 (starts at arr[1], ends at arr[4])

// 	fmt.Println(sum(10, 20, 30))
// 	fmt.Println(sum(slice...)) //ellipsis spread all the value of slice to variadic fn

// }

/*
Variadic - A Variadic fn is just a fn that accept a variable num of args of the same type
Go uses ...(ellipsis operator) for defining and expanding variadic arguments
*/

// func sum(number ...int) int { // Variadic fn
// 	total := 0
// 	for _, num := range number {
// 		total += num
// 	}
// 	return total
// }

/*
Append - The built-in append function is used to dynamically add elements to a slice:
If the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it.

append() is variadic, the following are all valid:

slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)

append() on anything other than itself is usually a BAD idea.
The append() function only creates a new array when there isn't any capacity left.
Again, to avoid bugs like this, you should always use the append function on the same slice the result is assigned to
*/

/*
Range - Go provides syntactic sugar to iterate easily over elements of a slice:

fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape
*/

//Slices practice:

// package main

// import (
// 	"fmt"
// 	"unicode"
// 	"unicode/utf8"
// )

// type Message interface {
// 	Type() string
// }

// type TextMessage struct {
// 	Sender  string
// 	Content string
// }

// func (tm TextMessage) Type() string {
// 	return "text"
// }

// type MediaMessage struct {
// 	Sender    string
// 	MediaType string
// 	Content   string
// }

// func (mm MediaMessage) Type() string {
// 	return "media"
// }

// type LinkMessage struct {
// 	Sender  string
// 	URL     string
// 	Content string
// }

// func (lm LinkMessage) Type() string {
// 	return "link"
// }

// func filterMessages(messages []Message, filterType string) []Message {
// 	var filteredMessages []Message

// 	for _, mess := range messages {
// 		if mess.Type() == filterType {
// 			filteredMessages = append(filteredMessages, mess)
// 		}
// 	}
// 	return filteredMessages
// }

// func main() {
// messages := []Message{
// 	TextMessage{"Alice", "Hello!"},
// 	MediaMessage{"Bob", "image", "photo.png"},
// 	LinkMessage{"Charlie", "https://example.com", "Check this out!"},
// 	TextMessage{"Dave", "Good morning!"},
// }

// textMessages := filterMessages(messages, "link")

// fmt.Println("Filtered Text Messages:")
// for _, msg := range textMessages {
// 	fmt.Println(msg)
// }

// fmt.Println(isValidPassword("helLo1"))               // true ✅ (Valid)
// fmt.Println(isValidPassword("hello1"))               // false ❌ (No uppercase)
// fmt.Println(isValidPassword("HELLO"))                // false ❌ (No digit)
// fmt.Println(isValidPassword("Hi123"))                // true ✅ (Valid)
// fmt.Println(isValidPassword("H"))                    // false ❌ (Too short)
// fmt.Println(isValidPassword("ThisIsALongPassword1")) // false ❌ (Too long)
//}

//Validate password

// At least 5 characters long but no more than 12 characters.
// Contains at least one uppercase letter.
// Contains at least one digit.

// func isValidPassword(password string) bool {
// 	len := utf8.RuneCountInString(password)

// 	if len < 5 || len > 12 {
// 		return false
// 	}

// 	hasUpper := false
// 	hasDigit := false

// 	for _, char := range password {
// 		if unicode.IsUpper(char) {
// 			hasUpper = true
// 		}
// 		if unicode.IsDigit(char) {
// 			hasDigit = true
// 		}
// 	}

// 	return hasUpper && hasDigit
// }

/*
maps:map keys may be of any type that is comparable.
eg: boolean, numeric, string, pointer, channel, and interface types, and structs or arrays

slices, maps, and functions hese types cannot be compared using ==, and can not be used as map keys.

good practice: instead of using nest maps use struct directly as a key

Insert - m[keys] = elem
Get - elem = m[keys]
Check if a key exist - elem,ok := m[key] - elem is the value and ok is boolean acc to the result
Delete - delete(m,key)
*/
// package main

// import (
// 	"fmt"
// )

// import "fmt"

// import (
// 	"errors"
// 	"fmt"
// )

// type User struct {
// 	name                 string
// 	phoneNumber          int
// 	scheduledForDeletion bool
// }

// func getUserMap(names []string, phoneNumber []int, scheduledForDeletion []bool) (map[string]User, error) {

// 	if len(names) != len(phoneNumber) || len(names) != len(scheduledForDeletion) {
// 		return nil, errors.New("invalid sizes")
// 	}

// 	userMap := make(map[string]User)

// 	for i := range names {
// 		userMap[names[i]] = User{
// 			name:                 names[i],
// 			phoneNumber:          phoneNumber[i],
// 			scheduledForDeletion: scheduledForDeletion[i],
// 		}
// 	}

// 	return userMap, nil
// }

// func deleteIfNecessary(users map[string]User, name string) (deleted bool, err error) {

// 	user, ok := users[name]

// 	if !ok {
// 		return false, errors.New("Not Found")
// 	}

// 	if !user.scheduledForDeletion {
// 		return false, nil
// 	}
// 	delete(users, name)
// 	return true, nil

// }

// func main() {
// names := []string{"Alice", "Bob", "Charlie"}
// phoneNumbers := []int{1234567890, 9876543210, 5556667777}
// scheduledForDeletion := []bool{false, true, false}

// result, err := getUserMap(names, phoneNumbers, scheduledForDeletion)
// if err != nil {
// 	fmt.Println("Error:", err)
// } else {
// 	fmt.Println("User Map:", result)
// }

// resultAfterDeletion, err2 := deleteIfNecessary(result, "Bob")

// if err2 != nil {
// 	fmt.Println("Error:", err)
// } else {
// 	fmt.Println("User Deletion:", resultAfterDeletion)
// 	fmt.Println(result)
// }

//another pr
// 	messagedUsers := []string{"alice", "bob", "alice", "charlie", "alice", "dave"}
// 	validUsers := map[string]int{"alice": 0, "bob": 0, "charlie": 0}

// 	getCounts(messagedUsers, validUsers)

// 	fmt.Println(validUsers)
// }

// func getCounts(messagedUsers []string, validUsers map[string]int) {

// 	for _, user := range messagedUsers {
// 		if _, ok := validUsers[user]; ok {
// 			validUsers[user]++
// 		}

//		}
//	}

// Nested maps

// import (
// 	"fmt"
// 	"strings"
// )

// func getNamesCount(names []string) map[rune]map[string]int {
// 	userNameCount := make(map[rune]map[string]int)
// 	for _, name := range names {
// 		firstChar := []rune(name)[0]

// 		if userNameCount[firstChar] == nil {
// 			userNameCount[firstChar] = make(map[string]int)
// 		}
// 		userNameCount[firstChar][name]++
// 	}
// 	return userNameCount
// }

// func countDistinctWords(messages []string) int {
// 	count := 0
// 	userMap := make(map[string]int)
// 	for _, message := range messages {
// 		name := strings.ToLower(message)
// 		_, ok := userMap[name]
// 		if !ok {
// 			count++
// 			userMap[name] = 1
// 		}
// 	}
// 	return count
// }

// func findSuggestedFriends(username string, friendships map[string][]string) []string {

// 	userFriendList, ok := friendships[username]

// 	if !ok {
// 		return nil
// 	}

// 	directFriend := make(map[string]bool)

// 	for _, friend := range userFriendList {
// 		directFriend[friend] = true
// 	}

// 	suggestedFriends := make(map[string]bool)

// 	for _, friend := range userFriendList {
// 		for _, friendsOfFriend := range friendships[friend] {
// 			if friendsOfFriend != username && !directFriend[friendsOfFriend] {
// 				suggestedFriends[friendsOfFriend] = true
// 			}
// 		}
// 	}

// 	if len(suggestedFriends) == 0 {
// 		return nil
// 	}

// 	var mulualFriend []string

// 	for friend := range suggestedFriends {
// 		mulualFriend = append(mulualFriend, friend)
// 	}
// 	return mulualFriend

// }

// func main() {
// names := []string{"billy", "billy", "bob", "joe", "apx13", "amit", "xlr8", "Amit", "XLr8"}
// result := countDistinctWords(names)

// for char, nameMap := range result {
// 	fmt.Printf("%c: %v\n", char, nameMap) // %c converts rune to char
// }
// 	friendships := map[string][]string{
// 		"Alice":   {"Bob", "Charlie"},
// 		"Bob":     {"Alice", "Charlie", "David"},
// 		"Charlie": {"Alice", "Bob", "David", "Eve"},
// 		"David":   {"Bob", "Charlie"},
// 		"Eve":     {"Charlie"},
// 	}

// 	res := findSuggestedFriends("Eve", friendships)
// 	fmt.Println(res)
// }

/*
Pointers:
A pointer is a variable that stores the memory address of another variable.

Empty pointer:
var p *int
it's zero val is nil. which means it doesn't point to any memory address. Empty pointers are also called "nil pointers".

If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a panic) that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying to dereference it.

Pointer receivers (*Type) allow methods to modify the struct.
Value receivers (Type) work on a copy, leaving the original unchanged.
Use pointer receivers for modifying methods or when dealing with large structs.

Which is typically more performant to pass to a function? (Especially for small amounts of data) - Value
*/

// func main() {

// 	myString := "hello"      // myString is just a string
// 	myStringPtr := &myString // myStringPtr is a pointer to myString's address

// 	fmt.Printf("value of myStringPtr: %v\n", *myStringPtr)

// 	text := "Hello, world! Welcome to the world of Go."
// 	newText := strings.ReplaceAll(text, "world", "universe")

// 	fmt.Println(newText)
// }

// type Analytics struct {
// 	MessagesTotal     int
// 	MessagesFailed    int
// 	MessagesSucceeded int
// }

// type Message struct {
// 	Recipient string
// 	Success   bool
// }

// func getMessageText(anal *Analytics, msg Message) {
// 	defer func() {
// 		(*anal).MessagesTotal++
// 	}()
// 	if msg.Success {
// 		(*anal).MessagesSucceeded++
// 	} else {
// 		(*anal).MessagesFailed++
// 	}
// }

// func main() {
// 	anal := &Analytics{}
// 	msg1 := Message{Recipient: "Alice", Success: true}
// 	msg2 := Message{Recipient: "Bob", Success: false}

// 	getMessageText(anal, msg1)
// 	getMessageText(anal, msg2)

// 	fmt.Printf("Total: %d, Succeeded: %d, Failed: %d\n", anal.MessagesTotal, anal.MessagesSucceeded, anal.MessagesFailed)
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type customer struct {
// 	id      int
// 	balance float64
// }

// type transactionType string

// const (
// 	transactionDeposit    transactionType = "deposit"
// 	transactionWithdrawal transactionType = "withdrawal"
// )

// type transaction struct {
// 	customerID      int
// 	amount          float64
// 	transactionType transactionType
// }

// func updateBalance(user *customer, trxn transaction) error {
// 	if trxn.transactionType != transactionDeposit && trxn.transactionType != transactionWithdrawal {
// 		return errors.New("unknown transaction type")
// 	}
// 	if trxn.transactionType == transactionDeposit {
// 		user.balance += trxn.amount
// 	}
// 	if trxn.transactionType == transactionWithdrawal {
// 		if user.balance < trxn.amount {
// 			return errors.New("insufficient funds")
// 		}
// 		user.balance -= trxn.amount
// 	}
// 	return nil
// }

/*
alternative:
better approach

func updateBalance(user *customer, trxn transaction) error {
	switch trxn.transactionType {
	case transactionDeposit:
		user.balance += trxn.amount
	case transactionWithdrawal:
		if user.balance < trxn.amount {
			return errors.New("insufficient funds")
		}
		user.balance -= trxn.amount
	default:
		return errors.New("unknown transaction type")
	}
	return nil
}

*/

// func main() {
// 	user := customer{
// 		id:      1,
// 		balance: 20.00,
// 	}

// 	trxn := transaction{
// 		customerID:      1,
// 		amount:          10,
// 		transactionType: "withdrawal",
// 	}

// 	updateBalance(&user, trxn)
// 	fmt.Println(user)
// }

/*
Package

naming - By convention, a package's name is the same as the last element of its import path. For instance, the math/rand package comprises files that begin with: package rand
Remember it's only a convention not a fixed way

Note:
 1. A directory of Go code can have at most one package
 2. Go programs are organized into packages. A package is a directory of Go code that's all compiled together (Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package)
 3. A repository contains one or more modules. A module is a collection of Go packages that are released together.
 4. A file named go.mod at the root of a project declares the module. It contains:
    The module path
    The version of the Go language your project requires
    Optionally, any external package dependencies your project has

To recap how packages and modules work in your project directory structure:

	You will have many git repositories on your machine (typically one per project).
	Each repository is typically a single module.
	Each repository contains one or more packages
	Each package consists of one or more Go source files in a single directory.
*/
/*
Concurrency:
In Go concurrency is handle through goroutine and channels.It achieves this using goroutines (lightweight threads) and channels (communication mechanisms).

Goroutine:
A goroutine is a function that runs concurrently with other functions. You can create one by using 'go' keyword. light weight threads.

Channels: (Do not communicate by sharing memory instead share memory by communicating)
Channels in go allow goroutine to communicate safely. They help avoide race conditions by providing a way to send and receive values between goroutine.

Buffered vs. Unbuffered → Buffered allows temporary storage, unbuffered blocks

Closing channels → Notifies receivers that no more values will be sent

Select → Waits for multiple channels to be ready

Blocking & Deadlocks : A deadlook is when a group of goroutine are all blocking so none of them can continue. Common bug in concurrent programming.
	Steps to follow to avoide Deadlocks and Blocking:
		1. Always have a goroutine waiting for the value before sending.
		2. A beffered Channels allow sending multiple val without an immediate receiver.
		3. Close channels properly (Blocking on receive - Use select with default, timeouts, or close channels properly.)
		4. Use worker pools to manage goroutines efficiently.
		5. Use struct{} for signaling when the value itself doesn’t matter.
		6. Always drain the channel properly in concurrent programs to avoid blocking.


	Blocking occurs when:

		A sender waits indefinitely because no goroutine is reading.
		A receiver waits indefinitely because no value is being sent.

Unbuffered channels force synchronization (send and receive must happen together)
Buffered - When to use: When you want the sender to keep going without waiting for the receiver,

When to close: When the sender’s done and you want receivers to know. Like ending a phone call.
Who closes - the sender closes not the receiver

Common notes: 1. : If main exits, all goroutines die. Use channels or a sync.WaitGroup to wait for them.
Range Over Channel (for msg := range ch)	When reading all values until the channel is closed

*/
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		ch <- "Hello from goroutine" //send data
// 	}()

// 	msg := <-ch //receive data
// 	fmt.Println(msg)
// 	close(ch)

// 	_, ok := <-ch

// 	if !ok {
// 		fmt.Println("Channel Closed")
// 	}

// 	/*
// 		 	defer func() {
// 				close(ch) //Closing signals that no more values will be sent.
// 			}()
// 	*/
// }

//Send email in batches.

// package main

// import "fmt"

// func addEmailsToQueue(emails []string) chan string {
// 	ch := make(chan string, len(emails))
// 	for _, e := range emails {
// 		ch <- e
// 	}
// 	close(ch)
// 	return ch
// }

// func main() {
// 	emails := []string{
// 		"amitkvs981@gmail.com",
// 		"amitkvs982@gmail.com",
// 		"amitprasad832104@gmail.com",
// 		"tldr@gmail.com",
// 	}
// 	emailsQueue := addEmailsToQueue(emails)

// 	go func() {
// 		for email := range emailsQueue {
// 			fmt.Println("Sending email to:", email)
// 		}
// 	}()

// 	fmt.Scanln()
// }

//Closing:

// package main

// import "fmt"

// func countReports(numSentCh chan int) int {
// 	count := 0
// 	for {
// 		val, ok := <-numSentCh

// 		if !ok {
// 			break
// 		}
// 		count += val
// 	}
// 	return count
// }

// func sendReports(numBatches int, ch chan int) {
// 	for i := 0; i < numBatches; i++ {
// 		numReports := i*23 + 32%17
// 		ch <- numReports
// 	}
// 	close(ch)
// }

// func main() {
// 	ch := make(chan int)
// 	go sendReports(4, ch)
// 	fmt.Println(countReports(ch))
// }

//Select:
// A select statement is used to listen to multiple channels at the same time. It is similar to a switch statement but for channels.
// select {
// case i, ok := <- chInts:
//     fmt.Println(i)
// case s, ok := <- chStrings:
//     fmt.Println(s)
// }

// package main

/*
Sleect Default case - the default case in the select smt executes immediately if no other channels has a value ready
	When you want non-blocking behavior.
	When you need to check a channel without waiting.
	When you want to avoid getting stuck in a select if no channels are ready.
*/

// func main() {
// time.Sleep(5000 * time.Millisecond) //wait execution for certain period of time
// fmt.Println("Waiting for 2 seconds...")
// <-time.After(5000 * time.Millisecond) // returns a channel immediately and sends a value after the specified time.
// fmt.Println("Time is up!")
// fmt.Println("Bro!")
// 	ticker := time.Tick(1 * time.Second) //This function returns a channel that sends a signal at regular intervals (like a ticker). It is useful for periodic tasks.

//	for i := 0; i < 5; i++ {
//		fmt.Println("ticking... ", ticker)
//		<-ticker
//	}
// }

// Types of channels:
// 1. Read Only channels
// 2. Write Only channels

// The name are itself pretty self explainatory let's look the example:
/*
func readCh(ch <-chan int) {
    // ch can only be read from
    // in this function
}

func writeCh(ch chan<- int) {
    // ch can only be written to
    // in this function
}
*/

/*
Nil Channel

var c chan string // c is nil
ch := make(chan string) // ch is initialize but not nil

Send - send to a nil channel Blocks forever c <- It's fucking blocks
Receive - Receive from a nil channles Block forever <-c

closed Channel
send - send to a closed channels panics
receive - receive to a closed channle return the zero value immediately.

*/

/*
sharing data - Mutex
more reads than writes - RWMutex
just need to pass data - Channels
*/

/*
Mutex
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}

	mu := &sync.RWMutex{}

	go writeLoop(m, mu)
	go readLoop(m, mu)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mu *sync.RWMutex) {
	for {
		for i := 0; i < 10; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
	}
}

func readLoop(m map[int]int, mu *sync.RWMutex) {
	for {
		mu.RLock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mu.RUnlock()
	}
}

/*
Generics :
func getLast[T any](s []T) T {
}
*/
