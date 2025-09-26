// for element add ke liye ek "append" function hoti h, jo ki h, inbuilt hoti h go me
package main

import "fmt"

func main() {

	var nums = make([]int, 2, 5)

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	fmt.Println(nums)
	// fmt.Println(cap(nums))
}
