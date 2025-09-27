// for example apko index ki help se iske andr element add krna h, to aap is tarike se kr skte h
package main

import "fmt"

func main() {
	var nums = make([]int, 2, 4)

	nums[0] = 3
	nums[0] = 5

	fmt.Println(nums)
	fmt.Println(cap(nums))

}
