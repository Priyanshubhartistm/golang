// functions in go lang
// function is a block of code that performs a specific task
// function is a reusable code block
// function is a way to organize code
// function is a way to modularize code
// function is a way to improve code readability
// function is a way to improve code maintainability
// function is a way to improve code reusability

package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

//2. multiple return values in function
func getlanguage() (string, string, string) {
	return "golang", "python", "javascript"
}

func main() {
	// result := add(3, 5)
	// println(result) // 8

	fmt.Println((getlanguage())) // golang python javascript

	// 3.  kise values ko agr ham use ni krte h, to ham _ (underscore) ka use kr skte h
	l1, l2, _ := getlanguage()
	fmt.Println(l1, l2) // golang python
}
