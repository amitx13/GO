package main

import "fmt"

const Secret = "preet is a good boy" // uppercase represents public constant, lowercase represents private constant

func main() {
	var userName string = "apx"
	fmt.Println("userName:", userName)
	fmt.Printf("Type %T \n", userName)

	var isUser bool = true
	fmt.Println("is a User:", isUser)
	fmt.Printf("Type %T \n", isUser)

	var UserAge int = 24 //"rune" can also be used for age as it is an alias for int32
	fmt.Println("Age of User:", UserAge)
	fmt.Printf("Type %T \n", UserAge)

	var userSalery float32 = 55.534232
	fmt.Println("userSalery of User:", userSalery)
	fmt.Printf("Type %T \n", userSalery)

	//default values
	var defaultName string
	fmt.Println("defaultName:", defaultName)
	fmt.Printf("Type %T \n", defaultName)
	var defaultAge int
	fmt.Println("defaultAge:", defaultAge)
	fmt.Printf("Type %T \n", defaultAge)
	var defaultSalery float64
	fmt.Println("defaultSalery:", defaultSalery)
	fmt.Printf("Type %T \n", defaultSalery)
	var defaultIsUser bool
	fmt.Println("defaultIsUser:", defaultIsUser)
	fmt.Printf("Type %T \n", defaultIsUser)

	//implicit declaration
	var implicitName = "apx"
	fmt.Println("implicitName:", implicitName)
	fmt.Printf("Type %T \n", implicitName)

	// short declaration || no var style || walrus operator ( := ) walrus can only be used inside methods not at package level or global level
	implicitAge := 24
	fmt.Println("implicitAge:", implicitAge)

	fmt.Println("Secret:", Secret)
	fmt.Printf("Type %T \n", Secret)
}
