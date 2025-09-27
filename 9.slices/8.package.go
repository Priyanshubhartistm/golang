// package in slices = do slices package k use krke , equal method  lga ke elements ko compair kr skte h.

package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}

	fmt.Println(slices.Equal(nums1, nums2))
	// dono slices equal h, to = true
	// dono slices equal ni h, to = false
}
