// In maps different keys and values then what

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["age"] = 30
	fmt.Println(m["phone"])
	// phone get krne ki kosis krenge, jo ki map ke andr exists hi ni krta h, to 0 values dega by default
}
