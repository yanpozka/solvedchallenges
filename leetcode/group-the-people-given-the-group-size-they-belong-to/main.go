package main

import "fmt"

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Println(groupThePeople([]int{2, 1, 3, 3, 3, 2}))
	fmt.Println(groupThePeople([]int{1, 1, 1, 1, 1}))
	fmt.Println(groupThePeople([]int{2, 2}))
}

func groupThePeople(groupSizes []int) [][]int {
	groups := map[int][]int{}
	for node, size := range groupSizes {
		groups[size] = append(groups[size], node)
	}

	var sol [][]int
	for size, nums := range groups {
		for ix := 0; ix < len(nums); ix += size {
			sol = append(sol, nums[ix:ix+size])
		}
	}
	return sol
}
