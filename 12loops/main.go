package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	frnd := []string{"Bihari", "chacha", "avhimanyu", "vky", "harshit", "rajan"}

	// for f := 0; f < len(frnd); f++ {
	// 	fmt.Println("Friend:", f, "is", frnd[f])
	// }

	for i, f := range frnd {
		fmt.Println("Friend:", i, "is", f)
	}

	rogueVal := 0

	for rogueVal < 5 {

		if rogueVal == 4 {
			goto apx
		}

		if rogueVal == 3 {
			fmt.Println("3, skipping this iteration")
			rogueVal++
			continue
			// fmt.Println("Rogue value is 3, breaking the loop")
			// break
		}
		fmt.Println("Rogue value is", rogueVal)
		rogueVal++
	}

apx:
	fmt.Println("This is apxx...")
}
