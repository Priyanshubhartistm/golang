// Pointers in go = variables that store the memory address of another variable
// pointers are used to store the address of a variable
// pointers are used to pass the address
// Pointers basicaaly use memory location ka address hota h.

package main

import "fmt"

func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum:", num) // 5
}

func main() {
	num := 10

	changeNum(num)

	fmt.Println("After changeNum:", num) // 10
	// yahan pr num ki value change ni hui, kyuki go lang me by default pass by value hota h, y fir aapko pointer ka use krna pdta h
	// pointer ka use krne k liye & use krte h
	// & = address of operator
	// * = dereference operator
}
