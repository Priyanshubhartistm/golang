// SWITCH = jab bhut jda conditions hote h, to ham switch k use krte h
// IF,ELSE = kam me if,else k kar lete h

package main

import "fmt"

func main() {
	// simple switch
	i := 3

	switch i {

	case 1:
		fmt.Println("one")

		// break keyword likhne ki jrurt ni h, oo automaticaly exit ho jata h.

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	case 4:
		fmt.Println("four")

		// default:
		// 	fmt.Println("other")

		// default case is optional
	}

}
