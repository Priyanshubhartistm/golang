// Another Slices in Go
package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}

	// if i add append
	nums = append(nums, 1)

	fmt.Println(nums)      // output = [1 2 3 4 5]
	fmt.Println(cap(nums)) // output = 5
	fmt.Println(len(nums)) // output = 5
}
