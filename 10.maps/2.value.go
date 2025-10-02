// In maps different keys and values then what

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["age"] = 30
	m["price"] = 50

	//fmt.Println(m["phone"])
	// phone get krne ki kosis krenge, jo ki map ke andr exists hi ni krta h, to 0 values dega by default
	// "zero" value diya bcz "int" type ki values h. agr boolean hota to "false" values deta.

	// fmt.Println(len(m)) // 1 length

	fmt.Println(len(m)) // 2 length
}
