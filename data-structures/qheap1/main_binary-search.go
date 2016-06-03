package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, action, val int
	fmt.Scan(&N)
	nums := make([]int, 0, N+1)

	for ; N > 0; N-- {
		fmt.Scan(&action)

		switch action {
		case 1:
			fmt.Scan(&val)
			if i := sort.SearchInts(nums, val); i == len(nums) {
				nums = append(nums, val)
			} else {
				nums = append(nums[:i], append([]int{val}, nums[i:]...)...)
			}

		case 2:
			fmt.Scan(&val)
			i := sort.SearchInts(nums, val)
			copy(nums[i:], nums[i+1:])
			nums[len(nums)-1] = 0
			nums = nums[:len(nums)-1]

		case 3:
			fmt.Println(nums[0])
		}
	}
}
