package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{2, 5, 5, 11}, 10))
}

func twoSum(nums []int, target int) []int {
	type pair struct {
		n, ix int
	}
	pairs := make([]*pair, len(nums))
	for ix, n := range nums {
		pairs[ix] = &pair{n: n, ix: ix}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].n < pairs[j].n
	})

	for ix, n := range nums {
		diff := target - n
		// fmt.Println("checking n:", n, " with diff:", diff)
		foundIx := sort.Search(len(pairs), func(i int) bool { return pairs[i].n >= diff })
		if foundIx == len(nums) {
			continue
		}

		p := pairs[foundIx]
		// fmt.Println("foundIx", foundIx, nums[p.ix], p.n)
		if nums[p.ix] == diff && ix != p.ix {
			return []int{ix, p.ix}
		}
	}
	return nil
}
