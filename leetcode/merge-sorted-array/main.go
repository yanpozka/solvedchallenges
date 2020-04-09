package main

import (
	"fmt"
	"sort"
)

func main() {
	{
		nums1 := []int{1, 2, 3, 0, 0, 0}
		merge(nums1, 3, []int{2, 5, 6}, 3)
		fmt.Println(nums1)
	}
	{
		nums1 := []int{1, 0}
		merge(nums1, 1, []int{2}, 1)
		fmt.Println(nums1)
	}
	{
		nums1 := []int{2, 0}
		merge(nums1, 1, []int{1}, 1)
		fmt.Println(nums1)
	}
	{
		nums1 := []int{4, 5, 6, 0, 0, 0}
		merge(nums1, 3, []int{1, 2, 3}, 3)
		fmt.Println(nums1)
	}
	{
		nums1 := []int{1, 2, 3, 0, 0, 0}
		merge(nums1, 3, []int{4, 5, 6}, 3)
		fmt.Println(nums1)
	}
	{
		nums1 := []int{4, 0, 0, 0, 0, 0}
		merge(nums1, 1, []int{1, 2, 3, 5, 6}, 5)
		fmt.Println(nums1)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for a := 0; a < n; a++ {
		ix := m
		for ; ix > 0 && nums2[a] < nums1[ix-1]; ix-- {
		}
		copy(nums1[ix+1:], nums1[ix:])
		nums1[ix] = nums2[a]
		m++
	}
}

func mergeFast(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
