package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 8))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 1))
	fmt.Println(fourSum([]int{0}, 0))
	fmt.Println(fourSum([]int{-1, 0, -5, -2, -2, -4, 0, 1, -2}, -9))
}

type quad struct {
	a, b, c, d int
}

func newQuad(a, b, c, d int) quad {
	ns := []int{a, b, c, d}
	sort.Ints(ns)
	return quad{ns[0], ns[1], ns[2], ns[3]}
}

type trio struct {
	a, b, c int
}

func newTrio(a, b, c int) trio {
	ns := []int{a, b, c}
	sort.Ints(ns)
	return trio{ns[0], ns[1], ns[2]}
}

func (t trio) find(ix int) bool {
	return t.a == ix || t.b == ix || t.c == ix
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	allEquals := true
	for ix, lim := 0, len(nums)-1; ix < lim; ix++ {
		if nums[ix] != nums[ix+1] {
			allEquals = false
			break
		}
	}
	if allEquals && nums[0]+nums[1]+nums[2]+nums[3] == target {
		return [][]int{{nums[0], nums[1], nums[2], nums[3]}}
	}

	threeSums := map[int][]trio{}
	for a := 0; a < len(nums)-1; a++ {
		for b := a + 1; b < len(nums); b++ {
			for c := b + 1; c < len(nums); c++ {
				sum := nums[a] + nums[b] + nums[c]
				threeSums[sum] = append(threeSums[sum], newTrio(a, b, c))
			}
		}
	}
	// fmt.Println(threeSums)
	mQ := map[quad]byte{}
	for ix := 0; ix < len(nums); ix++ {
		diff := target - nums[ix]
		trios, found := threeSums[diff]
		if !found {
			continue
		}
		// fmt.Println(diff, trios)
		for _, t := range trios {
			if !t.find(ix) {
				mQ[newQuad(nums[t.a], nums[t.b], nums[t.c], nums[ix])] = 1
			}
		}
	}

	var res [][]int
	for q := range mQ {
		res = append(res, []int{q.a, q.b, q.c, q.d})
	}
	return res
}
