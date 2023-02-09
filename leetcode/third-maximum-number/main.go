package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(thirdMax([]int{2, 2, 5}))
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{2, 1}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
	fmt.Println(thirdMax([]int{45, 2, 2, 3, 1}))
}

func thirdMax(nums []int) int {
	fm := nums[0]
	sm, thm := math.MinInt, math.MinInt

	for ix := 1; ix < len(nums); ix++ {
		n := nums[ix]
		if n > fm {
			thm = sm
			sm = fm
			fm = n
		} else if n != fm && n > sm {
			thm = sm
			sm = n
		} else if n != fm && n != sm && n > thm {
			thm = n
		}
	}

	// fmt.Println(fm, sm, thm)
	if thm == math.MinInt {
		return fm
	}

	return thm
}

func thirdMaxSort(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	thMax := nums[0]
	var flips int

	for ix := 1; ix < len(nums); ix++ {
		if thMax != nums[ix] {
			thMax = nums[ix]
			flips++

			if flips == 2 {
				return thMax
			}
		}
	}

	return nums[0]
}
