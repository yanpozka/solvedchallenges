package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}

func firstMissingPositive(nums []int) int {
	m := map[int]bool{}
	var or []int
	var min int = 1<<32 - 1
	for _, n := range nums {
		if n > 0 && !m[n] {
			or = append(or, n)
			m[n] = true
			if n < min {
				min = n
			}
		}
	}
	// fmt.Println(m)

	if min > 1 { // super early return
		return 1
	}

	sort.Ints(or)
	for ix, lim := 0, len(or)-1; ix < lim; ix++ {
		if or[ix+1]-or[ix] > 1 {
			return or[ix] + 1
		}
	}
	if or[len(or)-1] > 0 {
		return or[len(or)-1] + 1
	}
	return 1
}
