// for element add ke liye ek "append" function hoti h, jo ki h, inbuilt hoti h go me
package main

import "fmt"

func main() {

	// var nums = make([]int, 2, 5)
	var nums = make([]int, 0, 5) // normally yaha pe starting me 0 rakha jata h, q ki yaha p new slices add krni h, to yaha pe 0 rakha jata h, taki starting me jo do br 0 ata h, oo n aa jaye.
	// y ek empty slices yaha pe create ho jaye

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	fmt.Println(nums) // output = [0 0 1 2 3 4 5]
	// y jo output me 2 bar 0, 0 h, q ki int type h, isliya do br 0,0 h, agr boolean hota to 2 bar ; false, false aata

	fmt.Println(cap(nums)) // output = 10
	fmt.Println(len(nums))

}
