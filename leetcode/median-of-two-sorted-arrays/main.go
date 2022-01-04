package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 7}))
	fmt.Println(findMedianSortedArrays([]int{3}, []int{-2, -1}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 4}, []int{3}))
	fmt.Println(findMedianSortedArrays([]int{3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return med(nums2)
	}
	if len(nums2) == 0 {
		return med(nums1)
	}
	var nums []int
	if nums1[len(nums1)-1] <= nums2[0] {
		nums = nums1
		nums = append(nums, nums2...)
	} else if nums2[len(nums2)-1] <= nums1[0] {
		nums = nums2
		nums = append(nums, nums1...)
	} else {
		nums = make([]int, 0, len(nums1))
		var a, b int
		for a < len(nums1) && b < len(nums2) {
			if nums1[a] < nums2[b] {
				nums = append(nums, nums1[a])
				a++
			} else {
				nums = append(nums, nums2[b])
				b++
			}
		}
		if a < len(nums1) {
			nums = append(nums, nums1[a:]...)
		}
		if b < len(nums2) {
			nums = append(nums, nums2[b:]...)
		}
		//fmt.Println(a, b)
	}
	//fmt.Println(nums)

	return med(nums)
}

func med(nums []int) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	if len(nums) == 1 {
		return float64(nums[0])
	}
	h := len(nums) / 2
	if len(nums)%2 == 1 {
		return float64(nums[h])
	}
	return float64(nums[h]+nums[h-1]) / 2.0
}
