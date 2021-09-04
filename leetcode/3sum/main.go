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

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	type trio struct {
		a, b, c int
	}
	mTrio := map[trio]bool{}
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		pair := twoSum(nums, -n)
		if pair == nil {
			continue
		}
		if i == pair[0] || i == pair[1] {
			continue
		}
		// fmt.Println("n:", n, "pair:", pair)
		res := []int{n, nums[pair[0]], nums[pair[1]]}
		sort.Ints(res)
		t := trio{res[0], res[1], res[2]}
		if _, contains := mTrio[t]; !contains {
			result = append(result, res)
			mTrio[t] = true
		}
	}

	return result
}

func twoSum(nums []int, target int) []int {
	for ix, n := range nums {
		diff := target - n

		foundIx := sort.SearchInts(nums, diff)
		if foundIx == len(nums) {
			continue
		}
		if nums[foundIx] == diff && ix != foundIx {
			return []int{ix, foundIx}
		}
	}
	return nil
}
