// Variadic Functions in golang = Println(1,2,3,4,5,66, "hello") iske andr aap "n" number of parameters jo h, oo pass kr skte h

package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total

}

func main() {
	result := sum(3, 4, 5, 6, 88)

	fmt.Println(result)
}
