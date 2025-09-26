// Another Slices in Go
package main

import "fmt"

func main() {

	nums := []int{}

	fmt.Println(nums)      // output = [1 2 3 4 5]
	fmt.Println(cap(nums)) // output = 5
	fmt.Println(len(nums)) // output = 5
}
