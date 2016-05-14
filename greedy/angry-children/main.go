package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	nums := make([]int, N)

	for ix := range nums {
		fmt.Scan(&nums[ix])
	}

	sort.Ints(nums)

	min := 99999999999

	for ix, n := range nums {
		last := ix + K - 1

		if last >= N {
			break
		}
		if r := nums[last] - n; r < min {
			min = r
		}
	}

	fmt.Println(min)
}
