// type switch

package main

import "fmt"

func main() {

	whoAmI := func(i interface{}) {

		switch t := i.(type) {

		case int:
			fmt.Println("its an integer")

		case string:
			fmt.Println("its a string")

		case bool:
			fmt.Println("its a boolean")

		default:
			fmt.Println("other", t)

		}
	}

	whoAmI("golang")

}
