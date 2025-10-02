// operation in maps
// how to check equality in maps
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

// output true, because both maps are equal
// if both maps are not equal, then output false
// maps.Equal function use krke hum check kr skte h ki dono maps equal h ki ni
// maps.Equal function use krne ke liye "maps" package import krna padta h
