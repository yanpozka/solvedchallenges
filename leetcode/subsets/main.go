package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
}

func subsets(nums []int) [][]int {
	sets := [][]int{[]int{}}

	for c, lim := uint32(1), uint32(math.Pow(2, float64(len(nums)))); c < lim; c++ {
		var ns []int
		var ix int
		for mask := c; mask > 0; mask >>= 1 {
			if mask&1 == 1 {
				ns = append(ns, nums[ix])
			}
			ix++
		}
		sets = append(sets, ns)
	}

	return sets
}
