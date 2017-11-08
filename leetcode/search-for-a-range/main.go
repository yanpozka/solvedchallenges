package main

import (
	"fmt"
	"sort"
)

func main() {
	in := []int{5, 7, 7, 8, 8, 10}

	expected := []int{3, 4}
	result := searchRange(in, 8)

	fmt.Printf("%v equals to %v ?\n", result, expected)
}

func searchRange(nums []int, target int) []int {

	ix := sort.SearchInts(nums, target)
	if len(nums) == ix || nums[ix] != target {
		return []int{-1, -1}
	}

	return []int{ix, sort.SearchInts(nums, target+1) - 1}
}
