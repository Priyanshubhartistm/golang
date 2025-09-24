// capacity = in go maxium no. of elements can fit

package main

import "fmt"

func main() {

	var nums = make([]int, 20)

	fmt.Println(cap(nums))
}
