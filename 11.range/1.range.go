// Range :- A range is a sequence of numbers with a start and an end.
// The range function returns a slice of integers from start to end-1.
// The range function can also take a step parameter to specify the increment between numbers.

package main

import "fmt"

// iterating over data structures

func main() {
	nums := []int{2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

}
