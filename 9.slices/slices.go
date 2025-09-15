// Slices = this is a dynamic array(thats mean jab v aap slices declear krte ho to, koye v length nhi dete ho, aur jb v chahe usme elements add kr skte ho, remove kr skte ho, aur modify kr skte ho)
// iske size automatically adjust ho jata hai, jb v aap usme elements add ya remove krte ho)
// yeh ek reference type hai, jo underlying array ko point krta hai
// yeh bahut flexible aur powerful hota hai, kyunki yeh dynamically grow aur shrink ho sakta hai
// slices ko create krne ke liye, aap built-in make function ka use kr skte ho, ya existing array se slice create kr skte ho

package main

import "fmt"

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil
	var numbers []int
	println(numbers)

	// for nill check
	println(numbers == nil) // true

	// for check length
	println(len(numbers)) // 0

	// if you want your slices are not a null, then you should be use this , var nums = make([]int, 2)

	// make functions milti h, go ke andr oo ek inbuilt function h, usko ham use krenge uske andr ham dete h, [] bracket or type dete h, ki slices k type kya h,
	// () ise ke andr 2, 3 arguments dene padte h, jo ki h, initiale size, or size h, 2 ; mai chahta hu, y jo slices banegi uska size 2 hona chahiya.

	var nums = make([]int, 2)
	fmt.Println(nums)

}
