package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingNumber([]int{0}))
	fmt.Println(missingNumber([]int{1}))
	fmt.Println(missingNumber([]int{1, 2}))
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{0, 2}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func missingNumber(nums []int) int {
	var sum int

	for _, n := range nums {
		sum += n
	}

	n := len(nums)
	ts := (n * (n + 1)) / 2

	return ts - sum
}

func missingNumberSet(nums []int) int {
	checked := make([]bool, len(nums)+1)

	for _, n := range nums {
		checked[n] = true
	}

	for res, ok := range checked {
		if !ok {
			return res
		}
	}

	return 0
}
