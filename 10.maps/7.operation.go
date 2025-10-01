// operation in maps

package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{"price": 40, "phone": 100}

	m2 := map[string]int{"price": 40, "phone": 100}

	fmt.Println(maps.Equal(m1, m2)) // true

}
