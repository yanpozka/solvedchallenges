package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	const inf = -1000000
	var localMax, gMax int = 0, inf

	for _, n := range nums {
		localMax = max(n, n+localMax)
		if localMax > gMax {
			gMax = localMax
		}
	}
	return gMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
