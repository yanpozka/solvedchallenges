package main

import "fmt"

func main() {
	in := []int{1, 2, 3, 1}
	fmt.Printf("%v == %v ?\n", findPeakElement(in), 2)
}

func findPeakElement(nums []int) int {
	var prev, next int
	const minInt = -(1 << 63)

	for ix, n := range nums {
		prev, next = minInt, minInt

		if ix-1 >= 0 {
			prev = nums[ix-1]
		}
		if ix+1 < len(nums) {
			next = nums[ix+1]
		}

		if prev < n && n > next {
			return ix
		}
	}

	return len(nums) - 1
}
