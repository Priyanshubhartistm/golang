// variable checking in if else condition

package main

import "fmt"

func main() {

	if age := 20; age >= 18 {
		fmt.Println("You are an adult", age)

	} else if age >= 12 {
		fmt.Println("You are a teenager", age)
	}

}

// Go doesn't have ternary operator. agar apko kch check krna h,to normal if, else use krn hoga.

// ise wale version me ternary operator supported ni ho skta h, future me suppor aa jaye ho skta h..
