package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	fmt.Println(containsDuplicateMap([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for ix, lim := 0, len(nums)-1; ix < lim; ix++ {
		if nums[ix] == nums[ix+1] {
			return true
		}
	}

	return false
}

func containsDuplicateMap(nums []int) bool {
	checked := map[int]bool{}

	for _, n := range nums {
		if checked[n] {
			return true
		}

		checked[n] = true
	}

	return false
}
