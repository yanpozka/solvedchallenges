package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersect(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	for _, n := range nums1 {
		m1[n]++
	}
	m2 := map[int]int{}
	for _, n := range nums2 {
		m2[n]++
	}

	var result []int
	for n, c1 := range m1 {
		c2, ok := m2[n]
		if !ok {
			continue
		}

		occ := c1
		if c2 < occ {
			occ = c2
		}
		for ; occ > 0; occ-- {
			result = append(result, n)
		}
	}
	return result
}
