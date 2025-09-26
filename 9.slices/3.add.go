// for element add ke liye ek "append" function hoti h, jo ki h, inbuilt hoti h go me
package main

import "fmt"

func main() {

	var nums = make([]int, 2, 5)

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	fmt.Println(nums)      // output = [0 0 1 2 3 4 5]
	fmt.Println(cap(nums)) // output = 10
}
