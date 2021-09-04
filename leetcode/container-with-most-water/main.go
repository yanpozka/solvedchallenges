package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea([]int{1, 2, 1}))
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var max int

	for l, r := 0, len(height)-1; l < r; {
		h := int(math.Min(float64(height[l]), float64(height[r])))
		w := r - l
		if a := h * w; a > max {
			max = a
		}

		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
