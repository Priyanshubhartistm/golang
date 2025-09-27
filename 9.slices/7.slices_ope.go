// Slices operator = slices ke andr ki elements ko return krta h,

package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3}

	fmt.Println(nums[0:2]) // output = [1 2] becz 0 mtlb (0, index se start ho jayega kha tak 2, index tak)
	// to 0, index mtlb(1) and 2, index  mtlb (3)tak
	// iske beach ke jo v elements h, oo return ho jayegi(output)
	// but 2nd wali index ko exclude krta h, baki use phle wali ko return karega.

}
