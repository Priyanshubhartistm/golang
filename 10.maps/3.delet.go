// elements delete in maps

package main

import "fmt"

func main() {
	m := make([string]int)

	m["age"] = 30
	m["price"] = 50

	delete(m, "price") // price element delete krne k liye

	fmt.Println(m) // map[age:30]

}
