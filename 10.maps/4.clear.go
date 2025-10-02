// clear function in maps

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["age"] = 30
	m["price"] = 50

	clear(m)

	fmt.Println(m) // map empty ho gya

}
