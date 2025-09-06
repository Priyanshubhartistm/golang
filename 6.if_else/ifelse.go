package main

import "fmt"

func main() {
	// age := 28

	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// } else {

	// 	fmt.Println("You are a minor")

	// }

	age := 5

	if age >= 18 {
		fmt.Println("You are an adult")

	} else if age <= 12 {
		fmt.Println("You are a teenager")
	} else {

		fmt.Println("You are a kid")

	}

}
