// capacity = in go maxium no. of elements can fit , ise slice ki capacity kitni h, 2 iska matlb y ni h, ki 2 hi no. fit honge iske andr oo

package main

import "fmt"

func main() {

	// var nums = make([]int, 2)
	var nums = make([]int, 2, 5) // yaha pe hm initialse capacity v add kar skte h, but mai agr capacity ni dalungi to
	//  usi length k (yani usi size ki capacity yaha par aa jayegi jo ki, 2)aa jayegi
	// agr hame advance me pata h, ki mere slices ke andr kitne elements add hone wala h, to for example(5) to mai yaha par 3rd "parameters" de skta hu, jo ki h, initale capacity

	// yaha pe ek inbuilt method ati h, jo ki h,cap; or (nums) ko iske andr pass krenge or
	fmt.Println(cap(nums))
}
