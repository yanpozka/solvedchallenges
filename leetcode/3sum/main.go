package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{1, 2, -2, -1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}

type trio struct {
	a, b, c int
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	mTrio := map[trio]bool{}

	for f, lim := 0, len(nums)-1; f < lim; f++ {
		for s := f + 1; s < len(nums); s++ {
			third := -(nums[f] + nums[s])
			foundIx := sort.SearchInts(nums, third)
			if foundIx == f || foundIx == s || foundIx == len(nums) || nums[foundIx] != third {
				continue
			}
			mTrio[newOrderedTrio(nums[f], nums[s], nums[foundIx])] = true
		}
	}

	var result [][]int
	for t := range mTrio {
		result = append(result, []int{t.a, t.b, t.c})
	}
	return result
}

func newOrderedTrio(a, b, c int) trio {
	ns := []int{a, b, c}
	sort.Ints(ns)
	return trio{c: ns[2], b: ns[1], a: ns[0]}
}
