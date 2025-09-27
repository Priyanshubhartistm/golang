// Copy function in slices

package main

import "fmt"

func main() {

	var nums = make([]int, 0, 5)

	nums = append(nums, 2)
	var nums2 = make([]int, len(nums))

	// copy function
	copy(nums2, nums)

	fmt.Println(nums, nums2)

}
