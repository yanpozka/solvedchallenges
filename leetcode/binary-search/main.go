package main

import (
	"fmt"
	"sort"
)

func main() {
	l := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(l, 9))
	fmt.Println(searchStd(l, 9))

	{
		l := []int{-1, 0, 3, 5, 9, 12}
		fmt.Println(search(l, 2))
	}
	{
		l := []int{2}
		fmt.Println(search(l, 2))
	}

}

func search(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		h := l + (r-l)/2
		if nums[h] < target {
			l = h + 1
		} else if nums[h] > target {
			r = h
		} else {
			return h
		}
		// fmt.Println(l, r)
	}

	return -1
}

func searchStd(nums []int, target int) int {
	pos := sort.SearchInts(nums, target)
	if pos == len(nums) {
		return -1
	}
	if nums[pos] != target {
		return -1
	}
	return pos
}
