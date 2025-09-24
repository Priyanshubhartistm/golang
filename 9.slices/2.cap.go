// capacity = in go maxium no. of elements can fit , ise slice ki capacity kitni h, 2 iska matlb y ni h, ki 2 hi no. fit honge iske andr oo

package main

import "fmt"

func main() {

	var nums = make([]int, 20)

	// yaha pe ek inbuilt method ati h, jo ki h,cap; or (nums) ko iske andr pass krenge or
	fmt.Println(cap(nums))
}
