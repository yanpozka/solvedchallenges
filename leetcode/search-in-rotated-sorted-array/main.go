//
// https://play.golang.org/p/HzNcUUa2Ym
package main

import (
	"fmt"
	"sort"
)

func main() {
	in := []int{4, 5, 6, 7, 10, 1, 2}
	// in := []int{3}

	for _, n := range in {
		fmt.Println(n, search(in, n))
	}
}

func search(nums []int, target int) int {
	p := -1
	for ix, lim := 0, len(nums)-1; ix < lim; ix++ {
		if nums[ix] > nums[ix+1] {
			p = ix + 1
			break
		}
	}

	numbers := nums
	if p != -1 {
		numbers = append(nums[p:], nums[:p]...)
	}

	ix := sort.SearchInts(numbers, target)
	if ix == len(nums) || numbers[ix] != target {
		return -1
	}

	if p != -1 {
		lim := len(nums) - p
		if ix < lim {
			ix += p
		} else {
			ix = (ix+(p-lim)) % p
		}
	}
	return ix
}
