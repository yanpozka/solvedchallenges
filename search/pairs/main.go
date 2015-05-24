package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K, num, count int
	fmt.Scanf("%d %d", &N, &K)
	nums_dict := make(map[int]bool)
	nums_list := make([]int, N)

	for ix := 0; ix < N; ix++ {
		fmt.Scanf("%d ", &num)
		nums_dict[num] = true
		nums_list[ix] = num
	}
	sort.Ints(nums_list)

	for ix := 0; ix < N; ix++ {
		if nums_list[ix] >= K {
			var diff int = nums_list[ix] - K
			if nums_dict[diff] {
				count++
			}
		}
	}
	fmt.Println(count)
}
