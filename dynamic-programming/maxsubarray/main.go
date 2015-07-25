package main

import (
	"fmt"
)

var T, N, p int

func main() {
	fmt.Scanf("%d", &T)

	for ; T > 0; T-- {
		fmt.Scanf("%d", &N)
		nums := make([]int, N)

		for ix := 0; ix < N; ix++ {
			fmt.Scanf("%d", &nums[ix])
		}

		fmt.Printf("%d %d\n", maxSumContinuosArray(nums, N), maxSumNoContArray(nums, N))
	}
}

func maxSumContinuosArray(nm []int, N int) int {
	var sum, max_sum int = 0, 0
	var allNegatives bool = true

	for ix := 0; ix < N; ix++ {
		val := sum + nm[ix]
		if val > 0 {
			sum = val
		} else {
			sum = 0
		}

		if sum > max_sum {
			max_sum = sum
			allNegatives = false
		}
	}
	if allNegatives {
		var max int = nm[0]
		for ix := 1; ix < N; ix++ {
			if nm[ix] > max {
				max = nm[ix]
			}
		}
		return max
	}

	return max_sum
}

func maxSumNoContArray(nm []int, N int) int {
	var sum int = 0
	var allNegatives bool = true

	for ix := 0; ix < N; ix++ {
		if nm[ix] > 0 {
			sum += nm[ix]
			allNegatives = false
		}
	}
	if allNegatives {
		var max int = nm[0]
		for ix := 1; ix < N; ix++ {
			if nm[ix] > max {
				max = nm[ix]
			}
		}
		return max
	}

	return sum
}
